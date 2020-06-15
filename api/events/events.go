package events

import (
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

type AuthLoginEvent struct {
	UserNickname string `json:"user_nickname"`
}

func (e *AuthLoginEvent) String() string {
	return utils.PrettyStruct("Event", e)
}

type AuthLoginReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id,omitempty"`
}

func (e *AuthLoginReply) String() string {
	return utils.PrettyStruct("Reply", e)
}

type MessagesSendEvent struct {
	Text string `json:"text"`
}

func (e *MessagesSendEvent) String() string {
	return utils.PrettyStruct("Event", e)
}

type MessagesSendReply struct {
	Result    int    `json:"result"`
	MessageId string `json:"message_id,omitempty"`
}

func (e *MessagesSendReply) String() string {
	return utils.PrettyStruct("Reply", e)
}

type UsersGetAllEvent struct{}

func (e *UsersGetAllEvent) String() string {
	return utils.PrettyStruct("Event", e)
}

type UsersGetAllReply struct {
	Result int           `json:"result"`
	Users  []interface{} `json:"users,omitempty"` // Necessary!
}

func (e *UsersGetAllReply) String() string {
	return utils.PrettyStruct("Reply", e)
}

type UserNewUpdate struct {
	User *models.User `json:"user"`
}

func (e *UserNewUpdate) String() string {
	return utils.PrettyStruct("Event", e)
}

type UserDeletedUpdate struct {
	UserId string `json:"user_id"`
}

func (e *UserDeletedUpdate) String() string {
	return utils.PrettyStruct("Event", e)
}

type MessageNewUpdate struct {
	Message *models.Message `json:"message"`
}

func (e *MessageNewUpdate) String() string {
	return utils.PrettyStruct("Event", e)
}
