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
	Client *models.Client

	ClientPool *repository.Pool
	UserPool   *repository.Pool
}

func (h *GetAllEventHandler) NewEvent() network.Event {
	return &events.UsersGetAllEvent{}
}

func (h *GetAllEventHandler) ServeEvent(context.Context, network.Event) {
	if h.Client.UserId == "" {
		h.Client.Send("users.getAll", &events.UsersGetAllReply{Result: results.ErrInvalidSession})
		return
	}

	users := h.UserPool.GetAll("")
	h.Client.Send("users.getAll", &events.UsersGetAllReply{Users: users})
}
