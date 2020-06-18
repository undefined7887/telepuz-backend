package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/users/events"
)

type RemoveEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *RemoveEventHandler) NewEvent() network.Event {
	return nil
}

func (h *RemoveEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("users.removed", &events.Removed{UserId: h.Client.UserId})
	}

	h.ClientPool.Remove(h.Client.Id)
}
