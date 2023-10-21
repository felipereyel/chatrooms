package routes

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func initStaticRoutes(app *fiber.App) error {
	index := config.Configs.PublicDir + "/index.html"
	if !utils.CheckFileExists(index) {
		return errors.New("index.html not found")
	}

	app.Static("/", config.Configs.PublicDir)
	app.Static("*", index)

	return nil
}
