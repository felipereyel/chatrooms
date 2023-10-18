package routes

import (
	"chatrooms/gosrc/controllers"

	"github.com/gofiber/fiber/v2"
)

func initAuthRoutes(app fiber.Router, tc *controllers.UserController) {
	app.Get("/test", verifyAuth, testAuth)
	app.Post("/login", authLogin(tc), saveAuth)
	app.Post("/register", authRegister(tc), saveAuth)
}

func testAuth(c *fiber.Ctx) error {
	return c.SendString("authed")
}

func authLogin(tc *controllers.UserController) fiber.Handler {
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

func authRegister(tc *controllers.UserController) fiber.Handler {
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
