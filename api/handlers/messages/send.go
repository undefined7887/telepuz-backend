package messages

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/format"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/api/results"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/utils/rand"
)

type SendEventHandler struct {
	Client *models.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *SendEventHandler) NewEvent() network.Event {
	return &events.MessagesSendEvent{}
}

func (h *SendEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.MessagesSendEvent)

	if !h.checkEvent(event) {
		h.Client.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("messages.send", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

	message := &models.Message{
		Id:     rand.HexString(format.IdLength),
		UserId: h.Client.UserId,
		Text:   event.Text,
	}

	h.Client.Send("messages.send", &events.MessagesSendReply{MessageId: message.Id})
	h.Client.BroadcastSend("updates.message.new", &events.MessageNewUpdate{Message: message})
}

func (h *SendEventHandler) checkEvent(event *events.MessagesSendEvent) bool {
	return len(event.Text) > 0 && len(event.Text) < 6000
}
