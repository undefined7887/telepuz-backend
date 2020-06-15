package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type CloseEventHandler struct {
	Conn     network.Conn
	Session  *models.Session
	UserPool *repository.Pool
}

func (h *CloseEventHandler) NewEvent() network.Event {
	return nil
}

func (h *CloseEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Session.UserId != "" {
		h.UserPool.Remove(h.Session.UserId)
		h.Conn.BroadcastSend("updates.user.deleted", &events.UserDeletedUpdate{UserId: h.Session.UserId})
	}
}
