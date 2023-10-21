package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/database"

	"github.com/google/uuid"
)

type RoomController struct {
	dbRepo database.Database
}

func NewRoomController(dbRepo database.Database) *RoomController {
	return &RoomController{dbRepo}
}

func (tc *RoomController) ListRooms() ([]models.Room, error) {
	rooms, err := tc.dbRepo.ListRooms()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (tc *RoomController) CreateRoom(name string) (models.Room, error) {
	room := models.Room{
		Id:   uuid.New().String(),
		Name: name,
	}

	err := tc.dbRepo.CreateRoom(room)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func (tc *RoomController) GetRoom(roomId string) (models.Room, error) {
	room, err := tc.dbRepo.GetRoom(roomId)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}
