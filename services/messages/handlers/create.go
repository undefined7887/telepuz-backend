package handlers

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/messages/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

const (
	CreatePath  = "messages.create"
	CreatedPath = "messages.created"
)

type CreateMessage struct {
	MessageText string `json:"message_text"`
}

type CreateReply struct {
	Result    int    `json:"result"`
	MessageId string `json:"message_id,omitempty"`
}

type CreatedMessage struct {
	Message *models.Message `json:"message"`
}

type Create struct {
	*common.Client
	ClientPool *utils.Pool
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

	if h.UserId == "" {
		h.Client.SendMessage(CreatePath, &CreateReply{Result: common.ErrInvalidSessionResult})
		return
	}

	message := &models.Message{
		Id:     utils.RandHexString(common.IdLength),
		UserId: h.Client.UserId,
		Text:   event.MessageText,
	}

	h.SendMessage(CreatePath, &CreateReply{
		Result:    common.OkResult,
		MessageId: message.Id,
	})

	h.BroadcastMessage(CreatedPath, &CreatedMessage{Message: message})
}

func (h *Create) checkEvent(event *CreateMessage) bool {
	return len(event.MessageText) > 0 && len(event.MessageText) < 6000
}
