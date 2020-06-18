package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
)

type RemovedEvent struct {
	UserId string `json:"user_id"`
}

func (e *RemovedEvent) String() string {
	return log.PrettyStruct("Event", e)
}

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
		h.Client.BroadcastSend("users.removed", &RemovedEvent{UserId: h.Client.UserId})
	}

	h.ClientPool.Remove(h.Client.Id)
}
