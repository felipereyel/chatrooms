package controllers

import (
	"chatrooms/gosrc/repositories/database"
	"testing"
)

func TestRoomcontroller(t *testing.T) {
	dbRepo := database.FakeDatabaseRepo()
	roomController := NewRoomController(dbRepo)

	room, err := roomController.CreateRoom("test")
	if err != nil {
		t.Fatal(err)
	}

	room, err = roomController.GetRoom(room.Id)
	if err != nil {
		t.Fatal(err)
	}

	rooms, err := roomController.ListRooms()
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, r := range rooms {
		if r.Id == room.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("room %s not found in list", room.Id)
	}
}
