package handlers

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/users/format"
	"github.com/undefined7887/telepuz-backend/services/users/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

const (
	CreatePath  = "users.create"
	CreatedPath = "users.created"
)

type CreateMessage struct {
	UserNickname string `json:"user_nickname"`
}

type CreateReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id,omitempty"` // Necessary!
}

type CreatedMessage struct {
	User *models.User `json:"user"`
}

type Create struct {
	*common.Client

	ClientPool *utils.Pool
	UserPool   *utils.Pool
}

func (h *Create) NewMessage() interface{} {
	return &CreateMessage{}
}

func (h *Create) ServeMessage(messageInterface interface{}) {
	event := messageInterface.(*CreateMessage)

	if !h.checkEvent(event) {
		h.Client.SendMessage(CreatePath, &CreateReply{Result: common.ErrInvalidFormatResult})
		return
	}

	user := &models.User{
		Id:       utils.RandHexString(common.IdLength),
		Nickname: event.UserNickname,
		Status:   models.UserStatusOnline,
	}
	h.UserPool.Add(user.Id, user)

	if h.UserId != "" {
		h.UserPool.Remove(h.UserId)
		h.BroadcastMessage(RemovedPath, &RemovedMessage{UserId: h.UserId})
	}
	h.UserId = user.Id

	h.SendMessage(CreatePath, &CreateReply{
		Result: common.OkResult,
		UserId: user.Id,
	})

	h.BroadcastMessage(CreatedPath, &CreatedMessage{User: user})
}

func (h *Create) checkEvent(event *CreateMessage) bool {
	return format.UserNicknameRegexp.MatchString(event.UserNickname)
}
