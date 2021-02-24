package api

import (
	"fmt"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// CreateOneController is a controller that can be accessed through /api/new
func CreateOneController(c *fiber.Ctx) error {
	connection := CreateConnection()
	defer connection.Connection.Close()

	body := new(Post)
	body.IP = c.Hostname()

	fmt.Println(body)

	if err := c.BodyParser(body); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Bad JSON formatting",
			Success: false,
			Data:    nil,
		})
	}

	lastID, err := connection.savePost(body)

	if err := connection.insertTags(lastID, body.Tags); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
			Success: false,
			Data:    nil,
		})
	}

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
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
