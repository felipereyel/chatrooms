package routes

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/controllers"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) error {
	if config.Configs.JWTSecret == "" {
		return fmt.Errorf("[Init] JWTSecret is not set")
	}

	database, err := database.NewDatabaseRepo()
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	broker, err := broker.NewBrokerRepo()
	if err != nil {
		return fmt.Errorf("[Init] failed to get broker: %w", err)
	}

	apiGroup := app.Group("/_api")

	uc, err := controllers.NewUserController(database)
	if err != nil {
		return fmt.Errorf("[Init] failed to get user controller: %w", err)
	}
	usersGroup := apiGroup.Group("/users")
	initUsersRoutes(usersGroup, uc)

	rc := controllers.NewRoomController(database)
	roomsGroup := apiGroup.Group("/rooms")
	initRoomsRoutes(roomsGroup, rc)

	pc := controllers.NewPostsController(database, broker)
	postsGroup := apiGroup.Group("/rooms")
	initPostsRoutes(postsGroup, pc)

	return initStaticRoutes(app)
}
