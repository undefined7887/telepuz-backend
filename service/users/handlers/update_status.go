package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	models2 "github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
)

type UpdateStatusEvent struct {
	UserStatus int `json:"user_status"`
}

func (e *UpdateStatusEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type UpdateStatusReply struct {
	Result int `json:"result"`
}

func (e *UpdateStatusReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type StatusUpdatedEvent struct {
	UserId     string `json:"user_id"`
	UserStatus int    `json:"user_status"`
}

func (e *StatusUpdatedEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type UpdateStatusEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *UpdateStatusEventHandler) NewEvent() network.Event {
	return &UpdateStatusEvent{}
}

func (h *UpdateStatusEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*UpdateStatusEvent)

	if !h.checkEvent(event) {
		h.Client.Send("users.updateStatus", &UpdateStatusReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("users.updateStatus", &UpdateStatusReply{Result: results.ErrInvalidSession})
		return
	}

	user := h.UserPool.Get(h.Client.UserId).(*models2.User)
	user.Status = event.UserStatus
	h.UserPool.Add(user.Id, user)

	h.Client.Send("users.updateStatus", &UpdateStatusReply{Result: results.Ok})
	h.Client.BroadcastSend("user.statusUpdated", &StatusUpdatedEvent{UserId: user.Id, UserStatus: user.Status})
}

func (h *UpdateStatusEventHandler) checkEvent(event *UpdateStatusEvent) bool {
	return event.UserStatus > -1 && event.UserStatus < 3
}
