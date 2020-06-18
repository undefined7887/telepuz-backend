package users

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service/events"
	"github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
)

type UpdateStatusEventHandler struct {
	Client *models.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *UpdateStatusEventHandler) NewEvent() network.Event {
	return &events.UsersUpdateStatusEvent{}
}

func (h *UpdateStatusEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.UsersUpdateStatusEvent)

	if !h.checkEvent(event) {
		h.Client.Send("users.updateStatus", &events.UsersUpdateStatusReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("users.updateStatus", &events.UsersUpdateStatusReply{Result: results.ErrInvalidSession})
		return
	}

	user := h.UserPool.Get(h.Client.UserId).(*models.User)
	user.Status = event.Status
	h.UserPool.Add(user.Id, user)

	h.Client.Send("users.updateStatus", &events.UsersUpdateStatusReply{})
	h.Client.BroadcastSend("updates.user.newStatus", &events.UserNewStatusUpdate{UserId: user.Id, UserStatus: user.Status})
}

func (h *UpdateStatusEventHandler) checkEvent(event *events.UsersUpdateStatusEvent) bool {
	return event.Status > -1 && event.Status < 3
}
