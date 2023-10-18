package server

import (
	"chatrooms/gosrc/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) {
	app := fiber.New(fiber.Config{})
	app.Use(cors.New())

	app.Static("/", config.Configs.PublicDir)
	app.Static("*", config.Configs.PublicDir+"/index.html")

	if err := app.Listen(config.Configs.ServerAddress); err != nil {
		panic(err.Error())
	}
}
