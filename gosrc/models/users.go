package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
}
