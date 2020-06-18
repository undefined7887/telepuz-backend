package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/common/format"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	"github.com/undefined7887/telepuz-backend/service/messages/events"
	"github.com/undefined7887/telepuz-backend/service/messages/models"
	usersModels "github.com/undefined7887/telepuz-backend/service/users/models"
)

type SendEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *SendEventHandler) NewEvent() network.Event {
	return &events.SendEvent{}
}

func (h *SendEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.SendEvent)

	if !h.checkEvent(event) {
		h.Client.Send("messages.send", &events.SendReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("messages.send", &events.SendReply{Result: results.ErrInvalidSession})
		return
	}

	// We need to update status to online after message is sent
	user := h.UserPool.Get(h.Client.UserId).(*usersModels.User)
	user.Status = usersModels.UserStatusOnline
	h.UserPool.Add(user.Id, user)

	message := &models.Message{
		Id:     rand.HexString(format.IdLength),
		UserId: h.Client.UserId,
		Text:   event.MessageText,
	}

	h.Client.Send("messages.send", &events.SendReply{MessageId: message.Id})
	h.Client.BroadcastSend("updates.message.new", &events.NewUpdate{Message: message})
}

func (h *SendEventHandler) checkEvent(event *events.SendEvent) bool {
	return len(event.MessageText) > 0 && len(event.MessageText) < 6000
}
