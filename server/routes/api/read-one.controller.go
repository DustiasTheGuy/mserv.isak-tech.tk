package api

import "github.com/gofiber/fiber/v2"

// ReadOneController is a controller that can be accessed through /api/post/:ID
func ReadOneController(c *fiber.Ctx) error {
	return c.SendString(c.Params("ID"))
}
