package analytics

import (
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

func SaveNewRequestController(c *fiber.Ctx) error {

	return c.JSON(routes.HTTPResponse{
		Message: "",
		Success: true,
		Data:    nil,
	})
}

func GetAllRequestController(c *fiber.Ctx) error {
	return c.JSON(routes.HTTPResponse{
		Message: "",
		Success: true,
		Data:    nil,
	})
}

func GetSingleRequestController(c *fiber.Ctx) error {
	return c.JSON(routes.HTTPResponse{
		Message: "",
		Success: true,
		Data:    nil,
	})
}
