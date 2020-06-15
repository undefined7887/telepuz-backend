package websocket

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/vmihailenco/msgpack/v4"
	"time"
)

type conn struct {
	logger log.Logger
	inner  *websocket.Conn

	timeout  time.Duration
	handlers map[string]network.EventHandler
}

func (c *conn) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *conn) Handle(path string, handler network.EventHandler) {
	c.handlers[path] = handler
}

func (c *conn) Send(path string, event network.Event) {
	buffer := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buffer).UseJSONTag(true)

	if err := encoder.EncodeString(path); err != nil {
		c.logger.Fatal("Failed to write path: %s", err.Error())
	}

	if err := encoder.Encode(event); err != nil {
		c.logger.Fatal("Failed to write events: %s", err.Error())
	}

	if err := c.inner.WriteMessage(websocket.BinaryMessage, buffer.Bytes()); err != nil {
		c.logger.Warn("Failed to send message: %s", err.Error())
		return
	}

	c.logger.Info("Sent events \"%s\": %s", path, event)
}

func (c *conn) handleEvents() {
	for {
		_, msg, err := c.inner.ReadMessage()
		if err != nil {
			c.logger.Warn("Failed to receive message: %s", err.Error())
			return
		}

		decoder := msgpack.NewDecoder(bytes.NewBuffer(msg)).UseJSONTag(true)

		path, err := decoder.DecodeString()
		if err != nil {
			c.logger.Warn("Failed to read path: %s", err.Error())
			continue
		}

		handler := c.handlers[path]
		if handler == nil {
			c.logger.Warn("Failed to find handler for path: %s", path)
			continue
		}

		event := handler.NewEvent()
		if err := decoder.Decode(event); err != nil {
			c.logger.Warn("Failed to read events: %s", err.Error())
			continue
		}

		c.logger.Info("Received events \"%s\": %s", path, event)

		ctx, _ := context.WithTimeout(context.TODO(), c.timeout)
		go handler.ServeEvent(ctx, event)
	}
}

func NewConn(logger log.Logger, inner *websocket.Conn) network.Conn {
	logger = logger.WithPrefix(fmt.Sprintf("websocket-connection [%s]", inner.RemoteAddr()))

	conn := &conn{
		logger:   logger,
		inner:    inner,
		handlers: make(map[string]network.EventHandler),
	}

	go conn.handleEvents()
	return conn
}
