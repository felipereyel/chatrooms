package routes

import (
	"chatrooms/gosrc/utils"

	"github.com/gofiber/fiber/v2"
)

var cookieName = "chatrooms:jwt"
var headerName = "Authorization"

func verifyAuth(c *fiber.Ctx) error {
	headerJwt := c.Get(headerName)
	jwt := c.Cookies(cookieName, headerJwt)
	if jwt == "" {
		return fiber.ErrUnauthorized
	}

	id, err := utils.ParseJWT(jwt)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("id", id)
	return c.Next()
}

func saveAuth(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	jwt, exp, err := utils.GenerateJWT(id)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    cookieName,
		Value:   jwt,
		Expires: exp,
	})

	return c.SendStatus(fiber.StatusOK)
}
