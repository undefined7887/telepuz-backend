package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type listener struct {
	logger log.Logger
	inner  http.ServeMux

	conns   map[int]*conn
	handler network.ConnHandler
}

func (l *listener) Handle(handler network.ConnHandler) {
	l.handler = handler
}

func (l *listener) listen(addr string) {
	if err := http.ListenAndServe(addr, &l.inner); err != nil {
		l.logger.Fatal("Failed to listen for websocket connections: %s", err.Error())
	}
}

func (l *listener) handleConns(writer http.ResponseWriter, req *http.Request) {
	innerConn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		l.logger.Warn("Failed to accept innerConn: %s", err.Error())
		return
	}

	l.logger.Info("Accepted connection: %s", innerConn.RemoteAddr())
	conn := newConn(l.logger, l, innerConn)

	go l.handler.ServeConn(conn)
}

func NewListener(logger log.Logger, path, addr string) network.Listener {
	logger = logger.WithPrefix("websocket-listener")

	listener := &listener{logger: logger, conns: make(map[int]*conn)}

	listener.inner.HandleFunc(path, listener.handleConns)
	go listener.listen(addr)

	return listener
}
