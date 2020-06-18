package usersModels

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
}

const (
	UserStatusOffline = iota
	UserStatusOnline
	UserStatusTexting
)
