package messages

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/messages/handlers"
	"github.com/undefined7887/telepuz-backend/utils"
)

func NewService(client *common.Client, clientPool *utils.Pool) {
	client.OnMessage(handlers.CreatePath, &handlers.Create{
		Client:     client,
		ClientPool: clientPool,
	})
}
