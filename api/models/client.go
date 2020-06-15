package models

import (
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type Client struct {
	ClientPool *repository.Pool

	Id   string
	Conn network.Conn

	UserId string
}

func (c *Client) BroadcastOthersWithUserId(path string, event network.Event) {
	clients := c.ClientPool.GetAll(c.Id)

	for _, clientInterface := range clients {
		client := clientInterface.(*Client)

		if client.UserId != "" {
			client.Conn.Send(path, event)
		}
	}
}
