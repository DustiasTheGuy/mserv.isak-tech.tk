package api

import (
	"fmt"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// CreateOneController is a controller that can be accessed through /api/new
func CreateOneController(c *fiber.Ctx) error {
	connection := CreateConnection()
	body := new(Post)

	body.IP = c.Hostname()

	if err := c.BodyParser(body); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	lastID, err := connection.savePost(body)

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: "New Post Created",
		Success: true,
		Data:    lastID,
	})
}
