package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/results"
)

type GetEvent struct{}

func (e *GetEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type GetReply struct {
	Result int           `json:"result"`
	Users  []interface{} `json:"users,omitempty"` // Necessary!
}

func (e *GetReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type GetEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *GetEventHandler) NewEvent() network.Event {
	return &GetEvent{}
}

func (h *GetEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId == "" {
		h.Client.Send("users.get", &GetReply{Result: results.ErrInvalidSession})
		return
	}

	h.Client.Send("users.get", &GetReply{Result: results.Ok, Users: h.UserPool.GetAll("")})
}
