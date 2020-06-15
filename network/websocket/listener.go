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

	conns    []network.Conn
	handlers []network.ConnHandler
}

func (l *listener) Handle(handler network.ConnHandler) {
	l.handlers = append(l.handlers, handler)
}

func (l *listener) BroadcastSend(path string, event network.Event, excludeConn network.Conn) {
	for _, conn := range l.conns {
		if conn != excludeConn {
			conn.Send(path, event)
		}
	}
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

	conn := NewConn(l.logger, innerConn)
	l.conns = append(l.conns, conn)

	l.logger.Info("Accepted connection: %s", innerConn.RemoteAddr())

	for _, handler := range l.handlers {
		handler.ServeConn(conn)
	}
}

func NewListener(logger log.Logger, path, addr string) network.Listener {
	logger = logger.WithPrefix("websocket-listener")

	listener := &listener{
		logger:   logger,
		conns:    make([]network.Conn, 0),
		handlers: make([]network.ConnHandler, 0),
	}

	listener.inner.HandleFunc(path, listener.handleConns)
	go listener.listen(addr)

	return listener
}
