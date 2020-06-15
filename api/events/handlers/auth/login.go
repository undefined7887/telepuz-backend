package auth

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

type LoginEventHandler struct {
	network.Conn
	network.Listener
	*models.Session
	UserPool *repository.Pool
}

func (h *LoginEventHandler) NewEvent() network.Event {
	return &events.AuthLoginEvent{}
}

func (h *LoginEventHandler) ServeEvent(_ context.Context, eventInterface network.Event) {
	event := eventInterface.(*events.AuthLoginEvent)

	if !h.checkEvent(event) {
		h.Conn.Send("auth.login", &events.AuthLoginReply{Result: results.ErrInvalidFormat})
		return
	}

	user := &models.User{
		Id:       rand.HexString(format.IdLength),
		Nickname: event.Nickname,
	}

	h.UserPool.Add(user.Id, user)

	if h.Session.UserId != "" {
		h.UserPool.Remove(h.Session.UserId)
		h.Listener.BroadcastSend("updates.user.deleted", &events.UserDeletedUpdate{UserId: h.Session.UserId})
	}

	h.Conn.Send("auth.login", &events.AuthLoginReply{UserId: user.Id})
	h.Listener.BroadcastSend("updates.user.new", &events.UserNewUpdate{User: user})
}

func (h *LoginEventHandler) checkEvent(event *events.AuthLoginEvent) bool {
	return format.UserNicknameRegexp.MatchString(event.Nickname)
}
