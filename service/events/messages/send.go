package messages

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service/events"
	"github.com/undefined7887/telepuz-backend/service/format"
	"github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
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

	// We need to update status to online after message is sent
	user := h.UserPool.Get(h.Client.UserId).(*models.User)
	user.Status = models.UserStatusOnline
	h.UserPool.Add(user.Id, user)

	message := &models.Message{
		Id:     rand.HexString(format.IdLength),
		UserId: h.Client.UserId,
		Text:   event.MessageText,
	}

	h.Client.Send("messages.send", &events.MessagesSendReply{MessageId: message.Id})
	h.Client.BroadcastSend("updates.message.new", &events.MessageNewUpdate{Message: message})
}

func (h *SendEventHandler) checkEvent(event *events.MessagesSendEvent) bool {
	return len(event.MessageText) > 0 && len(event.MessageText) < 6000
}
