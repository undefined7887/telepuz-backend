package methods

import (
	"github.com/undefined7887/telepuz-backend/cache"
	"github.com/undefined7887/telepuz-backend/format"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/services/base/endpoint"
	"github.com/undefined7887/telepuz-backend/services/users/models"
)

type loginMethodRequestData struct {
	Nickname string `json:"nickname"`
}

type loginMethod struct {
	userPool *cache.Pool
}

func NewLoginMethod(userPool *cache.Pool) *loginMethod {
	return &loginMethod{userPool: userPool}
}

func (m *loginMethod) NewRequest() *endpoint.Request {
	return &endpoint.Request{Data: &loginMethodRequestData{}}
}

func (m *loginMethod) Call(req *endpoint.Request) (res *endpoint.Response) {
	reqData, res := req.Data.(*loginMethodRequestData), &endpoint.Response{Result: endpoint.Ok}

	if !m.checkRequestData(reqData) {
		res.Result = endpoint.ErrInvalidFormat
		return
	}

	user := &models.User{
		Id:       rand.Hex(format.IdLength),
		Nickname: reqData.Nickname,
	}

	// Adding user to global pool
	m.userPool.Add(user)

	return res
}

func (m *loginMethod) checkRequestData(reqData *loginMethodRequestData) bool {
	return format.UserNicknameRegexp.MatchString(reqData.Nickname)
}
