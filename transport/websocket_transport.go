package transport

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/vmihailenco/msgpack/v4"
	"net/http"
)

type websocketListener struct {
	logger log.Logger
	addr   string

	connHandler ConnHandler
}

func (l *websocketListener) Listen() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(writer, req, nil)
		if err != nil {
			l.logger.Warn("failed to accept connection: %s", err.Error())
			return
		}

		l.logger.Info("accepted connection %s", conn.RemoteAddr())
		l.connHandler.ServeConn(newWebsocketConn(l.logger, conn))
	})

	err := http.ListenAndServe(l.addr, nil)
	if err != nil {
		l.logger.Fatal("failed to listen %s: %s", l.addr, err.Error())
	}
}

func (l *websocketListener) OnConn(handler ConnHandler) {
	l.connHandler = handler
}

func NewWebsocketListener(logger log.Logger, addr string) Listener {
	return &websocketListener{
		addr:   addr,
		logger: logger.WithPrefix("listener"),
	}
}

type websocketConn struct {
	logger log.Logger
	conn   *websocket.Conn

	closeHandler    CloseHandler
	messageHandlers map[string]MessageHandler
}

func (c *websocketConn) SendMessage(path string, message interface{}) {
	messageBuffer := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(messageBuffer).UseJSONTag(true)

	if err := encoder.EncodeString(path); err != nil {
		c.logger.Fatal("failed to encode path: %s", err.Error())
	}

	if err := encoder.Encode(message); err != nil {
		c.logger.Fatal("failed to encode message: %s", err.Error())
	}

	if err := c.conn.WriteMessage(websocket.BinaryMessage, messageBuffer.Bytes()); err != nil {
		c.logger.Warn("failed to send data: %s", err.Error())
	}

	c.logger.Info("sent message (path: %s):\n%s", path, log.PrettyStruct("Message", message))
}

func (c *websocketConn) Close() {
	if err := c.conn.Close(); err != nil {
		c.logger.Warn("failed to close: %s", err.Error())
	}

	c.logger.Info("closed")

	if c.closeHandler != nil {
		c.logger.Warn("failed to find close handler")
		return
	}

	go c.closeHandler.ServeClose()
}

func (c *websocketConn) OnMessage(path string, handler MessageHandler) {
	c.messageHandlers[path] = handler
}

func (c *websocketConn) OnClose(handler CloseHandler) {
	c.closeHandler = handler
}

func (c *websocketConn) handleMessages() {
	for {
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			c.logger.Info("closed: %s", err.Error())

			if c.closeHandler == nil {
				c.logger.Warn("failed to find close handler")
				return
			}

			c.closeHandler.ServeClose()
			return
		}

		messageBuffer := bytes.NewBuffer(messageBytes)
		decoder := msgpack.NewDecoder(messageBuffer).UseJSONTag(true)

		path, err := decoder.DecodeString()
		if err != nil {
			c.logger.Warn("failed to decode path: %s", err.Error())
			continue
		}

		messageHandler := c.messageHandlers[path]
		if messageHandler == nil {
			c.logger.Warn("failed to find message handler (path: %s)", path)
			continue
		}

		message := messageHandler.NewMessage()
		if err := decoder.Decode(message); err != nil {
			c.logger.Warn("failed to decode message: %s", err.Error())
			continue
		}

		c.logger.Info("received message (path: %s):\n%s", path, log.PrettyStruct("Message", message))
		go messageHandler.ServeMessage(message)
	}
}

func NewWebsocketConn(logger log.Logger, addr string) Conn {
	logger = logger.WithPrefix(fmt.Sprintf("connection (addr: %s)", addr))

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+addr, nil)
	if err != nil {
		logger.Fatal("failed to connect: %s", err.Error())
	}

	return newWebsocketConn(logger, conn)
}

func newWebsocketConn(logger log.Logger, inner *websocket.Conn) Conn {
	// Important because this function can be used internally without NewWebsocketConn() call
	logger = logger.WithPrefix(fmt.Sprintf("connection (addr: %s)", inner.RemoteAddr()))

	conn := &websocketConn{
		logger:          logger,
		conn:            inner,
		messageHandlers: make(map[string]MessageHandler),
	}

	go conn.handleMessages()
	return conn
}
