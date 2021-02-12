package main

import (
	"paste/middleware"
	routes "paste/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	api := app.Group("/api", middleware.ApiMiddleware)
	api.Post("/new", routes.CreateOneController)   // add new post
	api.Get("/posts", routes.ReadManyController)   // get all posts
	api.Get("/post/:ID", routes.ReadOneController) // get one post

	app.Listen(":8082")
}
