package api

import (
	"log"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

func UpdateOneHandler(c *fiber.Ctx) error {
	var body Post
	connection := CreateConnection()
	defer connection.Connection.Close()

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Bad JSON Formatting",
			Success: false,
			Data:    nil,
		})
	}

	if err := connection.updateOne(&body); err != nil {
		log.Fatal(err)
	}

	return c.JSON(routes.HTTPResponse{
		Message: "It works",
		Success: true,
		Data:    nil,
	})
}
