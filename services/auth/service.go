package auth

import (
	"github.com/undefined7887/telepuz-backend/base"
	"github.com/undefined7887/telepuz-backend/base/endpoint"
	"github.com/undefined7887/telepuz-backend/cache"
	"github.com/undefined7887/telepuz-backend/services/auth/methods"
)

type authService struct {
	methods map[string]endpoint.Method
}

func NewService(userPool *cache.Pool) base.Service {
	service := &authService{}

	service.methods = map[string]endpoint.Method{
		"auth.login": methods.NewLoginMethod(userPool),
	}

	return service
}

func (s *authService) GetMethod(name string) endpoint.Method {
	return s.methods[name]
}