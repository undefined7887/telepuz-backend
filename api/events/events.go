package events

import (
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

type AuthLoginEvent struct {
	Nickname string `json:"nickname"`
}

func (e *AuthLoginEvent) String() string {
	return utils.PrettyStruct("Event [auth.login] \"", e)
}

type AuthLoginReply struct {
	Result int    `json:"results"`
	UserId string `json:"user_id"`
}

func (e *AuthLoginReply) String() string {
	return utils.PrettyStruct("Reply [auth.login]", e)
}

type MessagesSendEvent struct {
	models.Message
}

func (e *MessagesSendEvent) String() string {
	return utils.PrettyStruct("Event [messages.send]", e)
}

type MessagesSendReply struct {
	Result int `json:"results"`
}

func (e *MessagesSendReply) String() string {
	return utils.PrettyStruct("Reply [messages.send]", e)
}

type UsersGetAllEvent struct{}

func (e *UsersGetAllEvent) String() string {
	return utils.PrettyStruct("Event [users.getAll]", e)
}

type UsersGetAllReply struct {
	Result int            `json:"results"`
	Users  []*models.User `json:"users"`
}

func (e *UsersGetAllReply) String() string {
	return utils.PrettyStruct("Reply [users.getAll]", e)
}

type UserNewUpdate struct {
	User *models.User `json:"user"`
}

func (e *UserNewUpdate) String() string {
	return utils.PrettyStruct("Event [updates.user.new]", e)
}

type UserDeletedUpdate struct {
	UserId string `json:"user_id"`
}

func (e *UserDeletedUpdate) String() string {
	return utils.PrettyStruct("Event [updates.user.deleted]", e)
}

type MessageNewUpdate struct {
	Message *models.Message `json:"message"`
}

func (e *MessageNewUpdate) String() string {
	return utils.PrettyStruct("Event [updates.message.new]", e)
}
