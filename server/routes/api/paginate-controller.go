package api

import (
	"log"
	"paste/database"
	"paste/models/post"
	"paste/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PaginateController(c *fiber.Ctx) error {
	db, err := database.Connect("isak_tech_paste")
	defer db.Close()

	page, err := strconv.ParseInt(c.Params("PAGE"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	limit, err := strconv.ParseInt(c.Params("LIMIT"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if page < 0 {
		return c.JSON(routes.HTTPResponse{
			Message: "Cannot Execute Query",
			Success: false,
			Data:    nil,
		})

	} else if limit <= 0 {
		return c.JSON(routes.HTTPResponse{
			Message: "You want zero posts?",
			Success: false,
			Data:    nil,
		})

	}

	result, err := post.Paginate(limit, page*limit)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(routes.HTTPResponse{
		Message: "Works",
		Success: true,
		Data:    result,
	})
}
