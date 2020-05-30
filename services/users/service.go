package users

import (
	"github.com/undefined7887/telepuz-backend/cache"
	"github.com/undefined7887/telepuz-backend/services/base"
	"github.com/undefined7887/telepuz-backend/services/base/endpoint"
	"github.com/undefined7887/telepuz-backend/services/users/methods"
)

type usersService struct {
	methods map[string]endpoint.Method
}

func NewService(userPool *cache.Pool) base.Service {
	service := &usersService{}

	service.methods = map[string]endpoint.Method{
		"users.getAll": methods.NewGetAllMethod(userPool),
	}

	return service
}

func (s *usersService) GetMethod(name string) endpoint.Method {
	return s.methods[name]
}
