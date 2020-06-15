package updates

import (
	"github.com/undefined7887/telepuz-backend/api/models"
	"github.com/undefined7887/telepuz-backend/utils"
)

type UsersNewEvent struct {
	User *models.User `json:"user"`
}

func (e *UsersNewEvent) String() string {
	return utils.ToJSON("UserNewUpdate", e)
}
