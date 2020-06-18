package auth

import (
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	handlers "github.com/undefined7887/telepuz-backend/service/auth/handlers"
)

func NewService(client *service.Client, clientPool, userPool *repository.Pool) {
	client.Conn.Handle("auth.login", &handlers.LoginEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})
}
