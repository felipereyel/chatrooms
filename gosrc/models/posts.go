package models

type Post struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
}
