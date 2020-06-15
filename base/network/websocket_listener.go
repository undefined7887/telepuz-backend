package network

import (
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/base/log"
	"net/http"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type websocketListener struct {
	logger   log.Logger
	serveMux http.ServeMux

	handlers []ConnHandler
	conns    []Conn
}

func (l *websocketListener) Handle(handler ConnHandler) {
	l.handlers = append(l.handlers, handler)
}

func (l *websocketListener) BroadcastSend(path string, event Event) {
	for _, conn := range l.conns {
		conn.Send(path, event)
	}
}

func (l *websocketListener) listen(addr string) {
	if err := http.ListenAndServe(addr, &l.serveMux); err != nil {
		l.logger.Fatal("Failed to listen for websocket connections: %s", err.Error())
	}
}

func (l *websocketListener) handleConns(writer http.ResponseWriter, req *http.Request) {
	innerConn, err := websocketUpgrader.Upgrade(writer, req, nil)
	if err != nil {
		l.logger.Warn("Failed to accept innerConn: %s", err.Error())
		return
	}

	conn := NewWebsocketConn(l.logger, innerConn)
	l.conns = append(l.conns, conn)

	for _, handler := range l.handlers {
		handler(conn)
	}
}

func NewWebsocketListener(logger log.Logger, path, addr string) Listener {
	logger = logger.WithPrefix("websocket-listener")

	listener := &websocketListener{
		logger:   logger,
		handlers: make([]ConnHandler, 0),
		conns:    make([]Conn, 0),
	}

	listener.serveMux.HandleFunc(path, listener.handleConns)
	go listener.listen(addr)

	return listener
}
