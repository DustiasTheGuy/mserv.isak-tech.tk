package api

import (
	"log"
	"paste/database"
	"paste/models/post"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

func UpdateOneHandler(c *fiber.Ctx) error {
	var post *post.Post
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := c.BodyParser(&post); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Bad JSON Formatting",
			Success: false,
			Data:    nil,
		})
	}

	if err := post.UpdateOne(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(routes.HTTPResponse{
		Message: "It works",
		Success: true,
		Data:    nil,
	})
}
