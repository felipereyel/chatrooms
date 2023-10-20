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

	uc := controllers.NewUserController(database)
	usersGroup := apiGroup.Group("/users")
	initUsersRoutes(usersGroup, uc)

	rc := controllers.NewRoomController(database)
	roomsGroup := apiGroup.Group("/rooms")
	initRoomsRoutes(roomsGroup, rc)

	pc := controllers.NewPostsController(database)
	postsGroup := apiGroup.Group("/rooms")
	initPostsRoutes(postsGroup, pc)

	initStaticRoutes(app)
	return nil
}
