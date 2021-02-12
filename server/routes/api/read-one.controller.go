package api

import (
	"fmt"
	"mserv/database"
	"mserv/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ReadOneController is a controller that can be accessed through /api/post/:ID
func ReadOneController(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("ID"), 10, 32)

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	post, err := getPost(ID)

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
		Data:    post,
	})
}

// GetPosts from database
func getPost(ID int64) (Post, error) {
	var post Post

	db, err := database.Connect()

	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT * FROM posts WHERE ID=?", ID)

	if err := row.Scan(&post.ID, &post.Body, &post.Date); err != nil {
		return post, err
	}

	return post, nil
}
