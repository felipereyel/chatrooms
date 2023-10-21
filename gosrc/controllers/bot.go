package controllers

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"
	"chatrooms/gosrc/utils"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type BotController struct {
	dbRepo     database.Database
	brokerRepo broker.Broker
}

func NewBotController(dbRepo database.Database, brokerRepo broker.Broker) *BotController {
	return &BotController{dbRepo, brokerRepo}
}

func (tc *BotController) EnsureAccount() error {
	hashedPassword, err := utils.HashPassword(config.Configs.BotPassword)
	if err != nil {
		return err
	}

	user := models.User{
		Id:       uuid.New().String(),
		Username: config.Configs.BotUsername,
		Pass:     hashedPassword,
	}

	if err = tc.dbRepo.UserUpsert(user); err != nil {
		return err
	}

	return nil
}

func (bc *BotController) ListenAndAnswerCommands() error {
	consumer, err := bc.brokerRepo.ConsumeCommands()
	if err != nil {
		return err
	}
	defer consumer.Close()

	for msg := range consumer.Channel() {
		body := msg.Body
		var command models.CommandView
		if err := json.Unmarshal(body, &command); err != nil {
			// TODO handle bad message
			continue
		}

		response := models.PostView{
			RoomId:    command.RoomId,
			CreatedAt: utils.GetNow(),
			Id:        uuid.New().String(),
			Username:  config.Configs.BotUsername,
			Content:   fmt.Sprintf("Answer for: %s", command.Payload),
		}

		if err := bc.brokerRepo.PublishPost(command.RoomId, response); err != nil {
			// TODO handle bad message
			continue
		}

		msg.Ack(false)
	}

	return nil
}
