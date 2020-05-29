package models

type User struct {
	Id       string
	Nickname string
}

func (u *User) GetId() string {
	return u.Id
}
