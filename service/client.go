package service

import (
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type Client struct {
	Id     string
	UserId string

	Conn       network.Conn
	ClientPool *repository.Pool
}

func (c *Client) Send(path string, event network.Event) {
	c.Conn.Send(path, event)
}

func (c *Client) BroadcastSend(path string, event network.Event) {
	clients := c.ClientPool.GetAll(c.Id)

	for _, clientInterface := range clients {
		client := clientInterface.(*Client)

		if client.UserId != "" {
			client.Send(path, event)
		}
	}
}
