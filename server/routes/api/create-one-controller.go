package api

import (
	"log"
	"paste/database"
	"paste/models/post"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// CreateOneController is a controller that can be accessed through /api/new
func CreateOneController(c *fiber.Ctx) error {
	connection, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	var body *post.Post
	body.IP = c.Hostname()

	if err := c.BodyParser(body); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Bad JSON formatting",
			Success: false,
			Data:    nil,
		})
	}

	lastID, err := body.SavePost()

	if err := body.InsertTags(); err != nil {
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
