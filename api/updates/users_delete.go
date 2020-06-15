package updates

import (
	"github.com/undefined7887/telepuz-backend/utils"
)

type UsersDeleteEvent struct {
	UserId string `json:"user"`
}

func (e *UsersDeleteEvent) String() string {
	return utils.ToJSON("UserDeleteUpdate", e)
}
