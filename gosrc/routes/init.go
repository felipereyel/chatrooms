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

	apiGroup := app.Group("/_api")
	apiGroup.Get("/t", verifyAuth, func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	uc := controllers.NewUserController(database)
	authGroup := apiGroup.Group("/users")
	initUsersRoutes(authGroup, uc)

	initStaticRoutes(app)
	return nil
}
