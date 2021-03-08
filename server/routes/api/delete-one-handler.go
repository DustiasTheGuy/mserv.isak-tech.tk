package api

import (
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

func DeleteOneHandler(c *fiber.Ctx) error {
	connection := CreateConnection("isak_tech_paste")
	defer connection.Connection.Close()

	var post Post

	if err := c.BodyParser(&post); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Bad JSON Formatting",
			Success: false,
			Data:    nil,
		})
	}

	if err := connection.deleteOne(post.ID); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: "It works",
		Success: true,
		Data:    nil,
	})
}
