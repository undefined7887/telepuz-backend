package models

type Message struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
