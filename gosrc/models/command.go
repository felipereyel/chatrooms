package models

import (
	"regexp"
)

type CommandView struct {
	Id      string `json:"id"`
	Payload string `json:"payload"`
	RoomId  string `json:"room_id"`
}

func IsCommand(message string) bool {
	return message[0] == '/'
}

// Payload: /stock=stock_code
func (c *CommandView) IsValid() bool {
	regex := regexp.MustCompile(`^\/stock=[a-zA-Z0-9.]+$`)
	return regex.MatchString(c.Payload)
}

func (c *CommandView) FetchResponse() (string, error) {
	return "Answer for: " + c.Payload, nil
}
