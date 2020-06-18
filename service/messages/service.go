package messages

import (
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/messages/handlers"
)

func NewService(client *service.Client, clientPool, userPool *repository.Pool) {
	client.Conn.Handle("messages.create", &handlers.CreateEventHandler{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})
}
