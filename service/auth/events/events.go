package events

import "github.com/undefined7887/telepuz-backend/log"

type LoginEvent struct {
	UserNickname string `json:"user_nickname"`
}

func (e *LoginEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type LoginReply struct {
	Result int    `json:"result"`
	UserId string `json:"user_id,omitempty"`
}

func (e *LoginReply) String() string {
	return log.PrettyStruct("Reply", e)
}
