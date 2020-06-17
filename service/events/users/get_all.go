package users

import (
	"context"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service/events"
	"github.com/undefined7887/telepuz-backend/service/models"
	"github.com/undefined7887/telepuz-backend/service/results"
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

	h.Client.Send("users.getAll", &events.UsersGetAllReply{Users: h.UserPool.GetAll("")})
}
