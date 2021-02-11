package main

import (
	"mserv/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	apiGroup := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	apiGroup.Post("/new", api.NewPostHandler)
	apiGroup.Get("/posts", api.ReadPostsHandler)

	app.Listen(":8082")
}
