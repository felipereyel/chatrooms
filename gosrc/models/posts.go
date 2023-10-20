package models

type Post struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	RoomId  string `json:"room_id"`
}

type PostView struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}
