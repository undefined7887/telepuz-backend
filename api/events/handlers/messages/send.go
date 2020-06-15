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
	_ = eventInterface.(*events.MessagesSendEvent)

	if h.Session.UserId == "" {
		h.Conn.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

}
