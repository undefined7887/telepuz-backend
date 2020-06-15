package handlers

import (
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/base/repository"
)

type usersGetAllEvent struct{}

type usersGetAllEventReply struct {
	Users []models.User `json:"users"`
}

type getAllMethod struct {
	userPool *repository.Pool
}

func NewGetAllMethod(userPool *repository.Pool) endpoint.Method {
	return &getAllMethod{userPool: userPool}
}

func (m *getAllMethod) NewRequest() *endpoint.Request {
	return &endpoint.Request{Data: &getAllMethodRequestData{}}
}

func (m *getAllMethod) Call(_ *endpoint.Request) *endpoint.Response {
	res := &endpoint.Response{Result: endpoint.Ok}

	res.Data = &getAllMethodResponseData{
		Users: m.userPool.GetAll(),
	}

	return res
}
