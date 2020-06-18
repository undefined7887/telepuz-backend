package users

import (
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/users/handlers"
)

func NewService(client *service.Client, clientPool, userPool *repository.Pool) {
	client.Conn.Handle("users.create", &handlers.CreateEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.Conn.Handle("close", &handlers.RemoveEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.Conn.Handle("users.get", &handlers.GetEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.Conn.Handle("users.updateStatus", &handlers.UpdateStatusEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})
}
