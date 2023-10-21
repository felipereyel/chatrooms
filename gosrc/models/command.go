package models

import (
	"errors"
	"regexp"
)

type CommandView struct {
	Id      string `json:"id"`
	Payload string `json:"payload"`
	RoomId  string `json:"room_id"`
}

func IsCommand(message string) bool {
	return message != "" && message[0] == '/'
}

var ErrInvalidCommand = errors.New("invalid command")

func (c *CommandView) GetStockCode() (string, error) {
	regex := regexp.MustCompile(`^\/stock=([a-zA-Z0-9.]+)$`)
	matches := regex.FindStringSubmatch(c.Payload)
	if len(matches) != 2 {
		return "", ErrInvalidCommand
	}

	return matches[1], nil
}
