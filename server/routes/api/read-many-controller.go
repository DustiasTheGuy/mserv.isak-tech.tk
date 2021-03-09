package api

import (
	"fmt"
	"log"
	"paste/database"
	"paste/models/post"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// ReadManyController is a controller that can be accessed through /api/posts
func ReadManyController(c *fiber.Ctx) error {
	db, err := database.Connect("isak_tech_paste")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	posts, err := post.GetPosts()

	if err != nil {
		fmt.Println(err)
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
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
