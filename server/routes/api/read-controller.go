package api

import (
	"fmt"
	"mserv/routes"

	"github.com/gofiber/fiber/v2"
)

// ReadPostsHandler select * from posts
func ReadPostsHandler(c *fiber.Ctx) error {

	posts, err := GetPosts()

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
