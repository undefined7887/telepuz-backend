package methods

import (
	"github.com/undefined7887/telepuz-backend/cache"
	"github.com/undefined7887/telepuz-backend/services/base/endpoint"
)

type getAllMethodRequestData struct{}

type getAllMethodResponseData struct {
	Users []cache.Item `json:"users"`
}

type getAllMethod struct {
	userPool *cache.Pool
}

func NewGetAllMethod(userPool *cache.Pool) endpoint.Method {
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
