package database

import "chatrooms/gosrc/models"

// CRUDL
type Database interface {
	Close() error

	UserLogin(username string, password string) (models.User, error)
	UserRegister(user models.User) error
	UserUpsert(user models.User) error

	ListRooms() ([]models.Room, error)
	CreateRoom(room models.Room) error
	GetRoom(roomId string) (models.Room, error)

	ListPosts(roomId string) ([]models.PostView, error)
	CreatePost(post models.Post) (models.PostView, error)
}
