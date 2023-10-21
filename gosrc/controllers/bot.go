package controllers

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"
	"chatrooms/gosrc/repositories/stockapi"
	"chatrooms/gosrc/utils"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type BotController struct {
	botId      string
	dbRepo     database.Database
	brokerRepo broker.Broker
	stockApi   stockapi.StockApi
}

func NewBotController(dbRepo database.Database, brokerRepo broker.Broker, stockApi stockapi.StockApi) (*BotController, error) {
	hashedPassword, err := utils.HashPassword(config.Configs.BotPassword)
	if err != nil {
		return nil, err
	}

	var botId string
	botId, err = dbRepo.UserGetId(config.Configs.BotUsername)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		botId = uuid.New().String()
		botUser := models.User{
			Id:       botId,
			Username: config.Configs.BotUsername,
			Pass:     hashedPassword,
		}

		if err := dbRepo.UserRegister(botUser); err != nil {
			return nil, err
		}
	}

	return &BotController{botId, dbRepo, brokerRepo, stockApi}, nil
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
			// TODO handle internal bot errors
			continue
		}

		var responseContent string
		code, err := command.GetStockCode()
		if err != nil {
			responseContent = "[Error] Invalid command"
		} else {
			stockResponse, err := bc.stockApi.StockGet(code)
			if err != nil {
				responseContent = fmt.Sprintf("[Error] Processing failed: %s", err.Error())
			} else if stockResponse.Open == "N/D" {
				responseContent = fmt.Sprintf("%s quote is not available", stockResponse.Symbol)
			} else {
				responseContent = fmt.Sprintf("%s quote is $%s per share", stockResponse.Symbol, stockResponse.Open)
			}

		}

		post := models.Post{
			Id:      uuid.New().String(),
			UserId:  bc.botId,
			RoomId:  command.RoomId,
			Content: responseContent,
		}

		postview, err := bc.dbRepo.CreatePost(post)
		if err != nil {
			// TODO handle internal bot errors
			continue
		}

		if err := bc.brokerRepo.PublishPost(command.RoomId, postview); err != nil {
			// TODO handle internal bot errors
			continue
		}

		msg.Ack(false)
	}

	return nil
}
