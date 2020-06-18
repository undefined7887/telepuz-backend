package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/auth/events"
	"github.com/undefined7887/telepuz-backend/service/common/format"
	"github.com/undefined7887/telepuz-backend/service/common/results"
	usersEvents "github.com/undefined7887/telepuz-backend/service/users/events"
	usersModels "github.com/undefined7887/telepuz-backend/service/users/models"
)

type LoginEventHandler struct {
	Client *service.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *LoginEventHandler) NewEvent() network.Event {
	return &events.LoginEvent{}
}

func (h *LoginEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.LoginEvent)

	if !h.checkEvent(event) {
		h.Client.Send("auth.login", &events.LoginReply{Result: results.ErrInvalidFormat})
		return
	}

	user := &usersModels.User{
		Id:       rand.HexString(format.IdLength),
		Nickname: event.UserNickname,
		Status:   usersModels.UserStatusOnline,
	}
	h.UserPool.Add(user.Id, user)

	if h.Client.UserId != "" {
		h.UserPool.Remove(h.Client.UserId)
		h.Client.BroadcastSend("updates.user.deleted", &usersEvents.DeletedUpdate{UserId: h.Client.UserId})
	}

	h.Client.UserId = user.Id

	h.Client.Send("auth.login", &events.LoginReply{UserId: user.Id})
	h.Client.BroadcastSend("updates.user.new", &usersEvents.NewUpdate{User: user})
}

func (h *LoginEventHandler) checkEvent(event *events.LoginEvent) bool {
	return format.UserNicknameRegexp.MatchString(event.UserNickname)
}
