package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type CloseEventHandler struct {
	network.Conn
	network.Listener
	*models.Session
	UserPool *repository.Pool
}

func (h *CloseEventHandler) NewEvent() network.Event {
	return nil
}

func (h *CloseEventHandler) ServeEvent(context.Context, network.Event) {
	h.UserPool.Remove(h.UserId)

	if h.UserId != "" {
		h.BroadcastSend("updates.user.deleted", &events.UserDeletedUpdate{UserId: h.UserId}, h.Conn)
	}
}
