package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/database"
	"chatrooms/gosrc/utils"
	"errors"

	"github.com/google/uuid"
)

type UserController struct {
	dbRepo      database.Database
	botUsername string
}

func NewUserController(dbRepo database.Database, botUsername string) *UserController {
	return &UserController{dbRepo, botUsername}
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (tc *UserController) Login(req UserRequest) (string, error) {
	if req.Username == tc.botUsername {
		return "", errors.New("bot cannot login")
	}

	user, err := tc.dbRepo.UserLogin(req.Username, req.Password)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func (tc *UserController) Register(req UserRequest) (string, error) {
	if req.Username == tc.botUsername {
		return "", errors.New("bot cannot register")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Id:       uuid.New().String(),
		Username: req.Username,
		Pass:     hashedPassword,
	}

	err = tc.dbRepo.UserRegister(user)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}
