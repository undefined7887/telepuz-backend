package handlers

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api"
	"github.com/undefined7887/telepuz-backend/api/format"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/api/results"
	"github.com/undefined7887/telepuz-backend/api/updates"
	"github.com/undefined7887/telepuz-backend/base/network"
	"github.com/undefined7887/telepuz-backend/base/repository"
	"github.com/undefined7887/telepuz-backend/utils"
	"github.com/undefined7887/telepuz-backend/utils/rand"
)

type authLoginEvent struct {
	Nickname string `json:"nickname"`
}

func (e *authLoginEvent) String() string {
	return utils.ToJSON("AuthLoginEvent", e)
}

type authLoginReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id"`
}

func (e *authLoginReply) String() string {
	return utils.ToJSON("AuthLoginReply", e)
}

type authLoginHandler struct {
	listener network.Listener
	conn     network.Conn

	userPool *repository.Pool
}

func NewAuthLoginHandler(userPool *repository.Pool) *authLoginHandler {
	return &authLoginHandler{userPool: userPool}
}

func (h *authLoginHandler) NewEvent() network.Event {
	return &authLoginEvent{}
}

func (h *authLoginHandler) Handle(ctx context.Context, eventInterface network.Event) {
	event, reply := eventInterface.(*authLoginEvent), &authLoginReply{Result: results.Ok}

	if !h.checkEvent(event) {
		reply.Result = results.ErrInvalidFormat
		return
	}

	session := h.conn.Data().(*api.Session)
	if session.UserId != "" {
		h.userPool.Remove(session.UserId)
		h.listener.BroadcastSend("updates.user.delete", &updates.UsersDeleteEvent{UserId: session.UserId})
	}

	user := &models.User{
		Id:       rand.HexString(format.IdLength),
		Nickname: event.Nickname,
	}

	session.UserId = user.Id
	h.userPool.Add(user)

	reply.UserId = user.Id
	h.conn.Send("auth.login", reply)

	h.listener.BroadcastSend("updates.user.new", &updates.UsersNewEvent{User: user})
}

func (h *authLoginHandler) checkEvent(event *authLoginEvent) bool {
	return format.UserNicknameRegexp.MatchString(event.Nickname)
}
