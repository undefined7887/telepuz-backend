package events

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/service/messages/models"
)

type SendEvent struct {
	MessageText string `json:"message_text"`
}

func (e *SendEvent) String() string {
	return log.PrettyStruct("Event", e)
}

type SendReply struct {
	Result    int    `json:"result"`
	MessageId string `json:"message_id,omitempty"`
}

func (e *SendReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type NewUpdate struct {
	Message *models.Message `json:"message"`
}

func (e *NewUpdate) String() string {
	return log.PrettyStruct("Event", e)
}
