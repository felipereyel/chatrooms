package bot

import (
	"chatrooms/gosrc/controllers"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"
	"chatrooms/gosrc/repositories/stockapi"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	dbRepo, err := database.NewDatabaseRepo()
	if err != nil {
		panic(err.Error())
	}
	defer dbRepo.Close()

	brokerRepo, err := broker.NewBrokerRepo()
	if err != nil {
		panic(err.Error())
	}
	defer brokerRepo.Close()

	stockApi := stockapi.NewStockApi()

	bc, err := controllers.NewBotController(dbRepo, brokerRepo, stockApi)
	if err != nil {
		panic(err.Error())
	}

	if err := bc.ListenAndAnswerCommands(); err != nil {
		panic(err.Error())
	}
}
