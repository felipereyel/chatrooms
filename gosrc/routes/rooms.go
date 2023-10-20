package routes

import (
	"chatrooms/gosrc/controllers"

	"github.com/gofiber/fiber/v2"
)

func initRoomsRoutes(gp fiber.Router, rc *controllers.RoomController) {
	gp.Get("/", verifyAuth, getRooms(rc))
	gp.Post("/", verifyAuth, postRooms(rc))
	gp.Get("/:roomId", verifyAuth, getRoom(rc))
}

func postRooms(rc *controllers.RoomController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			Name string `json:"name"`
		}
		err := c.BodyParser(&body)
		if err != nil {
			// TODO handle bad request
			return err
		}

		room, err := rc.CreateRoom(body.Name)
		if err != nil {
			// TODO handle bad request
			return err
		}

		return c.JSON(room)
	}
}

func getRooms(rc *controllers.RoomController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rooms, err := rc.ListRooms()
		if err != nil {
			// TODO handle bad request
			return err
		}

		return c.JSON(rooms)
	}
}

func getRoom(rc *controllers.RoomController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		room, err := rc.GetRoom(c.Params("roomId"))
		if err != nil {
			// TODO handle bad request
			return err
		}

		return c.JSON(room)
	}
}
