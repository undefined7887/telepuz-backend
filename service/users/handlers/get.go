package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	"github.com/undefined7887/telepuz-backend/service/users/events"
)

type GetEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *GetEventHandler) NewEvent() network.Event {
	return &events.GetEvent{}
}

func (h *GetEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId == "" {
		h.Client.Send("users.get", &events.GetReply{Result: results.ErrInvalidSession})
		return
	}

	h.Client.Send("users.get", &events.GetReply{Users: h.UserPool.GetAll("")})
}
