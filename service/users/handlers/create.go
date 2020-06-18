package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/common/format"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	"github.com/undefined7887/telepuz-backend/service/users/events"
	"github.com/undefined7887/telepuz-backend/service/users/models"
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
		h.Client.Send("users.create", &events.CreateReply{Result: results.ErrInvalidFormat})
		return
	}

	user := &models.User{
		Id:       rand.HexString(format.IdLength),
		Nickname: event.UserNickname,
		Status:   models.UserStatusOnline,
	}
	h.UserPool.Add(user.Id, user)

	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("users.removed", &events.Removed{UserId: h.Client.UserId})
	}

	h.Client.UserId = user.Id

	h.Client.Send("users.create", &events.CreateReply{Result: results.Ok, UserId: user.Id})
	h.Client.BroadcastSend("users.created", &events.Created{User: user})
}

func (h *CreateEventHandler) checkEvent(event *events.Create) bool {
	return format.UserNicknameRegexp.MatchString(event.UserNickname)
}
