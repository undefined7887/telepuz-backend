package events

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/service/users/models"
)

type Create struct {
	UserNickname string `json:"user_nickname"`
}

func (e *Create) String() string {
	return log.PrettyStruct("Event", e)
}

type CreateReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id,omitempty"` // Necessary!
}

func (e *CreateReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type Get struct{}

func (e *Get) String() string {
	return log.PrettyStruct("Event", e)
}

type GetReply struct {
	Result int           `json:"result"`
	Users  []interface{} `json:"users,omitempty"` // Necessary!
}

func (e *GetReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type UpdateStatus struct {
	UserStatus int `json:"user_status"`
}

func (e *UpdateStatus) String() string {
	return log.PrettyStruct("Event", e)
}

type UpdateStatusReply struct {
	Result int `json:"result"`
}

func (e *UpdateStatusReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type Created struct {
	User *models.User `json:"user"`
}

func (e *Created) String() string {
	return log.PrettyStruct("Event", e)
}

type Removed struct {
	UserId string `json:"user_id"`
}

func (e *Removed) String() string {
	return log.PrettyStruct("Event", e)
}

type StatusUpdated struct {
	UserId     string `json:"user_id"`
	UserStatus int    `json:"user_status"`
}

func (e *StatusUpdated) String() string {
	return log.PrettyStruct("Event", e)
}
