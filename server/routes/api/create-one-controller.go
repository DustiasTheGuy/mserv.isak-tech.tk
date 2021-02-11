package api

import (
	"fmt"
	"mserv/database"
	"mserv/routes"

	"github.com/gofiber/fiber/v2"
)

// CreateOneController is a controller that can be accessed through /api/new
func CreateOneController(c *fiber.Ctx) error {
	body := new(Post)

	if err := c.BodyParser(body); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	lastID, err := savePost(body)

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: fmt.Sprint(err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: "HelloWorld",
		Success: true,
		Data:    lastID,
	})
}

// save a new post to the database
// first return value defaults to 0 if an error has occured
func savePost(body *Post) (int64, error) {
	db, err := database.Connect()

	if err != nil {
		return 0, err
	}

	sql := "INSERT INTO posts (body) VALUES (?)"
	res, err := db.Exec(sql, body.Body)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	defer db.Close()
	return lastID, nil
}
