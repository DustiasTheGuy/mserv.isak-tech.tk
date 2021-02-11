package index

import (
	"fmt"
	"mserv/routes"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Post is a struct for dealing with new Post
type Post struct {
	ID   uint      `json:"_id"`
	Body string    `json:"body"`
	Date time.Time `json:"date"`
}

// NewPostHandler saves a new post to the database
func NewPostHandler(c *fiber.Ctx) error {
	body := new(Post)

	if err := c.BodyParser(body); err != nil {
		fmt.Println(err)

		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
			Success: false,
			Data:    nil,
		})
	}

	fmt.Println(body)

	return c.JSON(routes.HTTPResponse{
		Message: "HelloWorld",
		Success: true,
		Data:    nil,
	})
}
