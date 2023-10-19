package routes

import (
	"chatrooms/gosrc/controllers"

	"github.com/gofiber/fiber/v2"
)

func initUsersRoutes(gp fiber.Router, tc *controllers.UserController) {
	gp.Post("/login", userLogin(tc), saveAuth)
	gp.Post("/register", userRegister(tc), saveAuth)
}

func userLogin(tc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req controllers.UserRequest
		err := c.BodyParser(&req)
		if err != nil {
			// TODO handle bad request
			return err
		}

		id, err := tc.Login(req)
		if err != nil {
			// TODO handle bad request
			return err
		}

		c.Locals("id", id)
		return c.Next()
	}
}

func userRegister(tc *controllers.UserController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req controllers.UserRequest
		err := c.BodyParser(&req)
		if err != nil {
			// TODO handle bad request
			return err
		}

		id, err := tc.Register(req)
		if err != nil {
			// TODO handle bad request
			return err
		}

		c.Locals("id", id)
		return c.Next()
	}
}
