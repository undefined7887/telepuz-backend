package network

import (
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/log"
	"net/http"
	"time"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type websocketTransport struct {
	logger log.Logger

	serveMux http.ServeMux
	timeout  time.Duration
}

func (w *websocketTransport) SetTimeout(timeout time.Duration) {
	w.timeout = timeout
}

func (w *websocketTransport) Handle(path string, handler Handler) {
	go w.handle(path, handler)
}

func (w *websocketTransport) handle(path string, handler Handler) {
	mux := &http.ServeMux{}

	mux.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
		conn, err := websocketUpgrader.Upgrade(writer, req, nil)
		if err != nil {
			w.logger.Warn("Failed to accept connection: %s", err.Error())
			return
		}

		go w.handleConn(conn, handler)
	})
}

func (w *websocketTransport) handleConn(conn *websocket.Conn, handler Handler) {
	for {
		msgType, msgBytes, err := conn.ReadMessage()
		if err != nil {

		}
	}
}

func (w *websocketTransport) HandleQueue(path, queue string, handler Handler) {
	panic("implement me")
}

func (w *websocketTransport) Send(path string, event Event) {
	panic("implement me")
}

func NewWebsocketsTransport(logger log.Logger) Transport {
	return &websocketTransport{logger: logger}
}
