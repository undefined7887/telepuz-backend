package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	"github.com/undefined7887/telepuz-backend/service/users/events"
)

type UpdateTypingEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *UpdateTypingEventHandler) NewEvent() network.Event {
	return &events.UpdateTypingEvent{}
}

func (h *UpdateTypingEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.UpdateTypingEvent)

	if h.Client.UserId == "" {
		h.Client.Send("users.updateTyping", &events.UpdateTypingReply{Result: results.ErrInvalidSession})
		return
	}

	h.Client.Send("users.updateTyping", &events.UpdateTypingReply{})
	h.Client.BroadcastSend("updates.user.newStatus", &events.NewTypingUpdate{UserId: h.Client.UserId, UserTyping: event.UserTyping})
}
