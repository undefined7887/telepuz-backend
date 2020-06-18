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

type CreateEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *CreateEventHandler) NewEvent() network.Event {
	return &events.Create{}
}

func (h *CreateEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.Create)

	if !h.checkEvent(event) {
		h.Client.Send("messages.create", &events.CreateReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("messages.create", &events.CreateReply{Result: results.ErrInvalidSession})
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

	h.Client.Send("messages.create", &events.CreateReply{Result: results.Ok, MessageId: message.Id})
	h.Client.BroadcastSend("messages.created", &events.Created{Message: message})
}

func (h *CreateEventHandler) checkEvent(event *events.Create) bool {
	return len(event.MessageText) > 0 && len(event.MessageText) < 6000
}
