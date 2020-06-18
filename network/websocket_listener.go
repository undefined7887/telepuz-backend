package network

import (
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type websocketListener struct {
	logger log.Logger
	inner  http.ServeMux

	handler ConnHandler
}

func (l *websocketListener) Handle(handler ConnHandler) {
	l.handler = handler
}

func (l *websocketListener) listen(addr string) {
	if err := http.ListenAndServe(addr, &l.inner); err != nil {
		l.logger.Fatal("Failed to listen for websocket connections: %s", err.Error())
	}
}

func (l *websocketListener) handleConns(writer http.ResponseWriter, req *http.Request) {
	innerConn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		l.logger.Warn("Failed to accept innerConn: %s", err.Error())
		return
	}

	l.logger.Info("Accepted connection: %s", innerConn.RemoteAddr())
	conn := newWebsocketConn(l.logger, innerConn)

	// This shouldn't block
	l.handler.ServeConn(conn)
}

func NewWebsocketListener(logger log.Logger, path, addr string) Listener {
	logger = logger.WithPrefix("websocket-websocketListener")

	listener := &websocketListener{logger: logger}

	listener.inner.HandleFunc(path, listener.handleConns)
	go listener.listen(addr)

	return listener
}
