package messages

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/format"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/api/results"
	"github.com/undefined7887/telepuz-backend/network"
)

type SendEventHandler struct {
	Conn    network.Conn
	Session *models.Session
}

func (h *SendEventHandler) NewEvent() network.Event {
	return &events.UsersGetAllEvent{}
}

func (h *SendEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.MessagesSendEvent)

	if !h.checkEvent(event) {
		h.Conn.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Session.UserId == "" {
		h.Conn.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

	h.Conn.Send("messages.send", &events.MessagesSendReply{})
	h.Conn.BroadcastSend("updates.message.new", &events.MessageNewUpdate{Message: &event.Message})
}

func (h *SendEventHandler) checkEvent(event *events.MessagesSendEvent) bool {
	return format.RegexpId.MatchString(event.Id) &&
		format.RegexpId.MatchString(event.UserId) &&
		len(event.Text) > 0 && len(event.Text) < 6000
}
