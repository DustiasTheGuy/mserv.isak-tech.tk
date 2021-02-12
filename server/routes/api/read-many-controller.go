package api

import (
	"fmt"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// ReadManyController is a controller that can be accessed through /api/posts
func ReadManyController(c *fiber.Ctx) error {
	connection := CreateConnection()
	posts, err := connection.getPosts()

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: nil,
		Success: true,
		Data:    posts,
	})
}
