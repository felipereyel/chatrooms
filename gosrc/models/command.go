package models

type CommandView struct {
	Id      string `json:"id"`
	Payload string `json:"payload"`
	RoomId  string `json:"room_id"`
}

func IsCommand(message string) bool {
	return message[0] == '/'
}
