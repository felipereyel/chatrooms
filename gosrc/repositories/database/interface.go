package database

import "chatrooms/gosrc/models"

// CRUDL
type Database interface {
	Close() error
	UserLogin(username string, password string) (models.User, error)
	UserRegister(user models.User) error
}
