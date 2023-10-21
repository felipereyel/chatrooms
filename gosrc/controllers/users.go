package controllers

import (
	"chatrooms/gosrc/config"
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

func NewUserController(dbRepo database.Database) (*UserController, error) {
	if config.Configs.BotUsername == "" {
		return nil, errors.New("BotUsername is not set")
	}

	return &UserController{dbRepo, config.Configs.BotUsername}, nil
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
