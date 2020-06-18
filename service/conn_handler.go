package service

import (
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service/events"
	"github.com/undefined7887/telepuz-backend/service/events/auth"
	"github.com/undefined7887/telepuz-backend/service/events/messages"
	"github.com/undefined7887/telepuz-backend/service/events/users"
	"github.com/undefined7887/telepuz-backend/service/format"
	"github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/utils/rand"
)

type connHandler struct {
	listener network.Listener

	clientPool *repository.Pool
	userPool   *repository.Pool
}

func (h *connHandler) ServeConn(conn network.Conn) {
	client := &models.Client{
		Id:         rand.HexString(format.IdLength),
		ClientPool: h.clientPool,
		Conn:       conn,
	}
	h.clientPool.Add(client.Id, client)

	conn.Handle("auth.login", &auth.LoginEventHandler{
		Client:     client,
		ClientPool: h.clientPool,
		UserPool:   h.userPool,
	})

	conn.Handle("users.getAll", &users.GetAllEventHandler{
		Client:     client,
		ClientPool: h.clientPool,
		UserPool:   h.userPool,
	})

	conn.Handle("users.updateStatus", &users.UpdateStatusEventHandler{
		Client:     client,
		ClientPool: h.clientPool,
		UserPool:   h.userPool,
	})

	conn.Handle("messages.send", &messages.SendEventHandler{
		Client:     client,
		ClientPool: h.clientPool,
		UserPool:   h.userPool,
	})

	conn.Handle("close", &events.CloseEventHandler{
		Client:     client,
		ClientPool: h.clientPool,
		UserPool:   h.userPool,
	})
}

func NewConnHandler(listener network.Listener, clientPool, userPool *repository.Pool) network.ConnHandler {
	return &connHandler{
		listener:   listener,
		clientPool: clientPool,
		userPool:   userPool,
	}
}
