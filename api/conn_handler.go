package api

import (
	"github.com/undefined7887/telepuz-backend/api/events/handlers"
	"github.com/undefined7887/telepuz-backend/api/events/handlers/auth"
	"github.com/undefined7887/telepuz-backend/api/events/handlers/messages"
	"github.com/undefined7887/telepuz-backend/api/events/handlers/users"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type connHandler struct {
	listener network.Listener
	userPool *repository.Pool
}

func (h *connHandler) ServeConn(conn network.Conn) {
	session := &models.Session{}

	conn.Handle("auth.login", &auth.LoginEventHandler{
		Conn:     conn,
		Listener: h.listener,
		Session:  session,
		UserPool: h.userPool,
	})

	conn.Handle("users.getAll", &users.GetAllEventHandler{
		Conn:     conn,
		Session:  session,
		UserPool: h.userPool,
	})

	conn.Handle("messages.send", &messages.SendEventHandler{
		Conn:     conn,
		Listener: h.listener,
		Session:  session,
	})

	conn.Handle("close", &handlers.CloseEventHandler{
		Conn:     conn,
		Listener: h.listener,
		Session:  session,
		UserPool: h.userPool,
	})
}

func NewConnHandler(listener network.Listener, userPool *repository.Pool) network.ConnHandler {
	return &connHandler{listener: listener, userPool: userPool}
}
