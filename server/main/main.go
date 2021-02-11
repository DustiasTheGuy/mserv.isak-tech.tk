package main

import (
	"mserv/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func apiMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	apiGroup := app.Group("/api", apiMiddleware)
	apiGroup.Post("/new", api.NewPostHandler)
	apiGroup.Get("/posts", api.ReadPostsHandler)

	app.Listen(":8082")
}
