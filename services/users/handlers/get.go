package handlers

import (
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/utils"
)

const GetPath = "users.get"

type GetMessage struct{}

type GetReply struct {
	Result int           `json:"result"`
	Users  []interface{} `json:"users,omitempty"` // Necessary!
}

type Get struct {
	*common.Client

	ClientPool *utils.Pool
	UserPool   *utils.Pool
}

func (h *Get) NewMessage() interface{} {
	return &GetMessage{}
}

func (h *Get) ServeMessage(interface{}) {
	if h.UserId == "" {
		h.SendMessage(GetPath, &GetReply{Result: common.ErrInvalidSessionResult})
		return
	}

	h.SendMessage(GetPath, &GetReply{
		Result: common.OkResult,
		Users:  h.UserPool.GetAll(""),
	})
}
