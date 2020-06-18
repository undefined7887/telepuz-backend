package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/format"
	"github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
)

type CreateEvent struct {
	MessageText string `json:"message_text"`
}

func (e *CreateEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type CreateReply struct {
	Result    int    `json:"result"`
	MessageId string `json:"message_id,omitempty"`
}

func (e *CreateReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type CreatedEvent struct {
	Message *models.Message `json:"message"`
}

func (e *CreatedEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type CreateEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *CreateEventHandler) NewEvent() network.Event {
	return &CreateEvent{}
}

func (h *CreateEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*CreateEvent)

	if !h.checkEvent(event) {
		h.Client.Send("messages.create", &CreateReply{Result: results.ErrInvalidFormat})
		return
	}

	if h.Client.UserId == "" {
		h.Client.Send("messages.create", &CreateReply{Result: results.ErrInvalidSession})
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

	h.Client.Send("messages.create", &CreateReply{Result: results.Ok, MessageId: message.Id})
	h.Client.BroadcastSend("messages.created", &CreatedEvent{Message: message})
}

func (h *CreateEventHandler) checkEvent(event *CreateEvent) bool {
	return len(event.MessageText) > 0 && len(event.MessageText) < 6000
}
