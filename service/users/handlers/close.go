package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/users/events"
)

type CloseEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *CloseEventHandler) NewEvent() network.Event {
	return nil
}

func (h *CloseEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("updates.user.deleted", &events.DeletedUpdate{UserId: h.Client.UserId})
	}

	h.ClientPool.Remove(h.Client.Id)
}
