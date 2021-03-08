package main

import (
	"paste/middleware"
	an "paste/routes/analytics"
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

	api := app.Group("/api", middleware.APIMiddleware)
	api.Post("/new", routes.CreateOneController)                 // add new post
	api.Get("/posts", routes.ReadManyController)                 // get all posts
	api.Get("/post/:ID", routes.ReadOneController)               // get one post
	api.Delete("/delete", routes.DeleteOneHandler)               // delete a single row
	api.Put("/update", routes.UpdateOneHandler)                  // update a single row
	api.Post("/sign-in", routes.SignInController)                // sign in
	api.Get("/paginate/:PAGE/:LIMIT", routes.PaginateController) // grab posts by page & limit

	analytics := api.Group("/an")
	analytics.Get("/", an.NewRequestController)

	app.Listen(":8082")
}
