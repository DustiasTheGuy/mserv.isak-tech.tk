package analytics

import (
	"paste/routes/api"

	"github.com/gofiber/fiber/v2"
)

func NewRequestController(c *fiber.Ctx) error {
	connection := api.CreateConnection("isak_tech_analytics")
	defer connection.Connection.Close()

	return nil
}
