package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)

	c.SendStatus(500)
	return c.JSON(fiber.Map{"error": err.Error()})
}
