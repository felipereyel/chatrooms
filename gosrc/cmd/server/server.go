package server

import (
	"chatrooms/gosrc/cmd/migrate"
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) {
	if config.Configs.AutoMigrate {
		migrate.Up(cmd, args)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler:  routes.ErrorHandler,
		StrictRouting: true,
	})

	app.Use(cors.New())
	routes.Init(app)

	if err := app.Listen(config.Configs.ServerAddress); err != nil {
		panic(err.Error())
	}
}
