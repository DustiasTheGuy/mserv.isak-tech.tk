package api

import (
	"fmt"
	"mserv/database"
	"mserv/routes"

	"github.com/gofiber/fiber/v2"
)

// NewPostHandler saves a new post to the database
func NewPostHandler(c *fiber.Ctx) error {
	body := new(Post)

	if err := c.BodyParser(body); err != nil {

		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	db, err := database.Connect()

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	insert, err := db.Query(fmt.Sprintf("INSERT INTO posts(body) VALUES('%s')", body.Body))

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	db.Close()
	defer insert.Close()

	return c.JSON(routes.HTTPResponse{
		Message: "HelloWorld",
		Success: true,
		Data:    nil,
	})
}
