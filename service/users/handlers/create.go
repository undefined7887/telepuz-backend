package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/format"
	models2 "github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
)

type CreateEvent struct {
	UserNickname string `json:"user_nickname"`
}

func (e *CreateEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type CreateReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id,omitempty"` // Necessary!
}

func (e *CreateReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type CreatedEvent struct {
	User *models2.User `json:"user"`
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
		h.Client.Send("users.create", &CreateReply{Result: results.ErrInvalidFormat})
		return
	}

	user := &models2.User{
		Id:       rand.HexString(format.IdLength),
		Nickname: event.UserNickname,
		Status:   models2.UserStatusOnline,
	}
	h.UserPool.Add(user.Id, user)

	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("users.removed", &RemovedEvent{UserId: h.Client.UserId})
	}

	h.Client.UserId = user.Id

	h.Client.Send("users.create", &CreateReply{Result: results.Ok, UserId: user.Id})
	h.Client.BroadcastSend("users.created", &CreatedEvent{User: user})
}

func (h *CreateEventHandler) checkEvent(event *CreateEvent) bool {
	return format.UserNicknameRegexp.MatchString(event.UserNickname)
}
