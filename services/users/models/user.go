package models

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func (u *User) GetId() string {
	return u.Id
}
