package api

import (
	"fmt"
	"paste/database"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

// ReadManyController is a controller that can be accessed through /api/posts
func ReadManyController(c *fiber.Ctx) error {
	posts, err := getPosts()

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

// GetPosts from database
func getPosts() ([]Post, error) {
	var Posts []Post
	db, err := database.Connect()

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post

		// for each row, scan the result into our tag composite object
		err = rows.Scan(&post.ID, &post.Body, &post.Date)

		if err != nil {
			return nil, err
		}

		Posts = append(Posts, post)
	}

	defer db.Close()

	return Posts, nil
}
