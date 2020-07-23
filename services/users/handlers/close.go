package handlers

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/utils"
)

const RemovedPath = "users.removed"

type RemovedMessage struct {
	UserId string `json:"user_id"`
}

type Close struct {
	*common.Client

	ClientPool *utils.Pool
	UserPool   *utils.Pool
}

func (h *Close) ServeClose() {
	if h.Client.UserId != "" {
		h.BroadcastMessage(RemovedPath, &RemovedMessage{UserId: h.UserId})
		h.UserPool.Remove(h.UserId)
	}

	h.ClientPool.Remove(h.Id)
}
