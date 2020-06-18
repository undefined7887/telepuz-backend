package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	"github.com/undefined7887/telepuz-backend/service/users/events"
	models "github.com/undefined7887/telepuz-backend/service/users/models"
)

type UpdateStatusEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *UpdateStatusEventHandler) NewEvent() network.Event {
	return &events.UpdateStatusEvent{}
}

func (h *UpdateStatusEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.UpdateStatusEvent)

	if !h.checkEvent(event) {
		h.Client.Send("users.updateStatus", &events.UpdateStatusReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("users.updateStatus", &events.UpdateStatusReply{Result: results.ErrInvalidSession})
		return
	}

	user := h.UserPool.Get(h.Client.UserId).(*models.User)
	user.Status = event.UserStatus
	h.UserPool.Add(user.Id, user)

	h.Client.Send("users.updateStatus", &events.UpdateStatusReply{})
	h.Client.BroadcastSend("updates.user.newStatus", &events.NewStatusUpdate{UserId: user.Id, UserStatus: user.Status})
}

func (h *UpdateStatusEventHandler) checkEvent(event *events.UpdateStatusEvent) bool {
	return event.UserStatus > -1 && event.UserStatus < 3
}
