package events

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/service/users/models"
)

type GetEvent struct{}

func (e *GetEvent) String() string {
	return log.PrettyStruct("Reply", e)
}

type GetReply struct {
	Result int           `json:"result"`
	Users  []interface{} `json:"users,omitempty"` // Necessary!
}

func (e *GetReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type UpdateStatusEvent struct {
	UserStatus int `json:"user_status"`
}

func (e *UpdateStatusEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type UpdateStatusReply struct {
	Result int `json:"result"`
}

func (e *UpdateStatusReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type NewUpdate struct {
	User *models.User `json:"user"`
}

func (e *NewUpdate) String() string {
	return log.PrettyStruct("Event", e)
}

type DeletedUpdate struct {
	UserId string `json:"user_id"`
}

func (e *DeletedUpdate) String() string {
	return log.PrettyStruct("Event", e)
}

type NewStatusUpdate struct {
	UserId     string `json:"user_id"`
	UserStatus int    `json:"user_status"`
}

func (e *NewStatusUpdate) String() string {
	return log.PrettyStruct("Event", e)
}
