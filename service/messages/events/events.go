package events

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/service/messages/models"
)

type Create struct {
	MessageText string `json:"message_text"`
}

func (e *Create) String() string {
	return log.PrettyStruct("Event", e)
}

type CreateReply struct {
	Result    int    `json:"result"`
	MessageId string `json:"message_id,omitempty"`
}

func (e *CreateReply) String() string {
	return log.PrettyStruct("Reply", e)
}

type Created struct {
	Message *models.Message `json:"message"`
}

func (e *Created) String() string {
	return log.PrettyStruct("Event", e)
}
