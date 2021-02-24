package api

import (
	"log"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

const (
	password = "hello"
)

func SignInController(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
	}

	if body["password"] != password {
		return c.JSON(routes.HTTPResponse{
			Message: "Denied",
			Success: false,
			Data:    nil,
		})

	} else if body["password"] == password {
		return c.JSON(routes.HTTPResponse{
			Message: "Approved",
			Success: true,
			Data:    nil,
		})

	}

	return c.JSON(routes.HTTPResponse{
		Message: "Denied",
		Success: false,
		Data:    nil,
	})
}
