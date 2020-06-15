package users

import (
	"context"
	"github.com/undefined7887/telepuz-backend/api/events"
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/api/results"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
)

type GetAllEventHandler struct {
	network.Conn
	*models.Session
	UserPool *repository.Pool
}

func (h *GetAllEventHandler) NewEvent() network.Event {
	return &events.UsersGetAllEvent{}
}

func (h *GetAllEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Session.UserId == "" {
		h.Conn.Send("users.getAll", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

	users := h.UserPool.GetAll()
	h.Conn.Send("users.getAll", &events.UsersGetAllReply{Users: users.([]*models.User)})
}
