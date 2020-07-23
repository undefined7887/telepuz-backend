package common

import (
	"github.com/undefined7887/telepuz-backend/transport"
	"github.com/undefined7887/telepuz-backend/utils"
)

type Client struct {
	Id     string
	UserId string

	conn       transport.Conn
	clientPool *utils.Pool
}

func (c *Client) SendMessage(path string, message interface{}) {
	c.conn.SendMessage(path, message)
}

func (c *Client) BroadcastMessage(path string, message interface{}) {
	clients := c.clientPool.GetAll(c.Id)

	for _, clientInterface := range clients {
		client := clientInterface.(*Client)

		if client.UserId != "" {
			client.SendMessage(path, message)
		}
	}
}

func (c *Client) OnMessage(path string, handler transport.MessageHandler) {
	c.conn.OnMessage(path, handler)
}

func (c *Client) OnClose(handler transport.CloseHandler) {
	c.conn.OnClose(handler)
}

func NewClient(conn transport.Conn, clientPool *utils.Pool) *Client {
	client := &Client{
		Id:         utils.RandHexString(IdLength),
		conn:       conn,
		clientPool: clientPool,
	}

	clientPool.Add(client.Id, client)
	return client
}
