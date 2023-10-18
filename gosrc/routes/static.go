package routes

import (
	"chatrooms/gosrc/config"

	"github.com/gofiber/fiber/v2"
)

func initStaticRoutes(app *fiber.App) {
	app.Static("/", config.Configs.PublicDir)
	app.Static("*", config.Configs.PublicDir+"/index.html")
}
