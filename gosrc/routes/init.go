package routes

import (
	"chatrooms/gosrc/controllers"
	"chatrooms/gosrc/repositories/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) error {
	database, err := database.NewDatabaseRepo()
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	uc := controllers.NewUserController(database)
	authApp := app.Group("/_auth")
	initAuthRoutes(authApp, uc)

	initStaticRoutes(app)
	return nil
}
