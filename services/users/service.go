package users

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/users/handlers"
	"github.com/undefined7887/telepuz-backend/utils"
)

func NewService(client *common.Client, clientPool, userPool *utils.Pool) {
	client.OnMessage(handlers.CreatePath, &handlers.Create{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.OnMessage(handlers.GetPath, &handlers.Get{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.OnMessage(handlers.UpdateStatusPath, &handlers.UpdateStatus{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})

	client.OnClose(&handlers.Close{
		Client:     client,
		ClientPool: clientPool,
		UserPool:   userPool,
	})
}
