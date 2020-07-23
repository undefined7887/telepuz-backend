package handlers

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/users/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

const (
	UpdateStatusPath  = "users.updateStatus"
	StatusUpdatedPath = "users.statusUpdated"
)

type UpdateStatusMessage struct {
	UserStatus int `json:"user_status"`
}

type UpdateStatusReply struct {
	Result int `json:"result"`
}

type StatusUpdatedMessage struct {
	UserId     string `json:"user_id"`
	UserStatus int    `json:"user_status"`
}

type UpdateStatus struct {
	*common.Client

	ClientPool *utils.Pool
	UserPool   *utils.Pool
}

func (h *UpdateStatus) NewMessage() interface{} {
	return &UpdateStatusMessage{}
}

func (h *UpdateStatus) ServeMessage(messageInterface interface{}) {
	event := messageInterface.(*UpdateStatusMessage)

	if !h.checkEvent(event) {
		h.SendMessage(UpdateStatusPath, &UpdateStatusReply{Result: common.ErrInvalidFormatResult})
		return
	}

	if h.UserId == "" {
		h.SendMessage(UpdateStatusPath, &UpdateStatusReply{Result: common.ErrInvalidSessionResult})
		return
	}

	user := h.UserPool.Get(h.UserId).(*models.User)
	user.Status = event.UserStatus

	h.UserPool.Add(user.Id, user)
	h.SendMessage(UpdateStatusPath, &UpdateStatusReply{Result: common.OkResult})

	h.BroadcastMessage(StatusUpdatedPath, &StatusUpdatedMessage{
		UserId:     user.Id,
		UserStatus: user.Status,
	})
}

func (h *UpdateStatus) checkEvent(event *UpdateStatusMessage) bool {
	return event.UserStatus > -1 && event.UserStatus < 3
}
