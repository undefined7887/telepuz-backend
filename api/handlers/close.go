package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type CloseEventHandler struct {
	Client *models.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *CloseEventHandler) NewEvent() network.Event {
	return nil
}

func (h *CloseEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("updates.user.deleted", &events.UserDeletedUpdate{UserId: h.Client.UserId})
	}

	h.ClientPool.Remove(h.Client.Id)
}
