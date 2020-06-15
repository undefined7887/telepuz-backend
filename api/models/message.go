package models

type Message struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}

func (u *Message) GetId() string {
	return u.Id
}
