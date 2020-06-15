package messages

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/api/results"
	"github.com/undefined7887/telepuz-backend/network"
)

type SendEventHandler struct {
	network.Conn
	network.Listener
	Session *models.Session
}

func (h *SendEventHandler) NewEvent() network.Event {
	return &events.UsersGetAllEvent{}
}

func (h *SendEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.MessagesSendEvent)

	if !h.checkEvent(event) {
		h.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Session.UserId == "" {
		h.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

	h.Send("messages.send", &events.MessagesSendReply{})
	h.BroadcastSend("messages.send", &events.MessageNewUpdate{Message: &event.Message})
}

func (h *SendEventHandler) checkEvent(event *events.MessagesSendEvent) bool {
	return true
}
