package main

import (
	"paste/middleware"
	"paste/routes/analytics"
	"paste/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Static("/", "./public")
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	apiRouter := app.Group("/api", middleware.APIMiddleware)
	apiRouter.Post("/new", api.CreateOneController)                 // add new post
	apiRouter.Get("/posts", api.ReadManyController)                 // get all posts
	apiRouter.Get("/post/:ID", api.ReadOneController)               // get one post
	apiRouter.Delete("/delete", api.DeleteOneHandler)               // delete a single row
	apiRouter.Put("/update", api.UpdateOneHandler)                  // update a single row
	apiRouter.Get("/paginate/:PAGE/:LIMIT", api.PaginateController) // grab posts by page & limit

	analyticsRouter := app.Group("/analytics")
	analyticsRouter.Post("/create", analytics.SaveNewRequestController)
	analyticsRouter.Get("/load/load_all", analytics.GetAllRequestController)
	analyticsRouter.Get("/load/:id", analytics.GetSingleRequestController)

	app.Listen(":8082")
}
