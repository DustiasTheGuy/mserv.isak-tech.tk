package analytics

import (
	"fmt"
	"paste/models/request"
	"paste/routes"

	"github.com/gofiber/fiber/v2"
)

func SaveNewRequestController(c *fiber.Ctx) error {
	var request *request.Request

	if err := c.BodyParser(&request); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Unable to parse data",
			Success: false,
			Data:    nil,
		})
	}

	request.IP = c.IP()

	if err := request.SaveRequest(); err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Unable to save data",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: "Everything went fine",
		Success: true,
		Data:    nil,
	})
}

func GetAllRequestController(c *fiber.Ctx) error {
	requests := request.GetAllRequests()

	return c.JSON(routes.HTTPResponse{
		Message: "",
		Success: true,
		Data:    requests,
	})
}

func GetSingleRequestController(c *fiber.Ctx) error {
	fmt.Printf("Get %s\n", c.Params("id"))
	return c.JSON(routes.HTTPResponse{
		Message: "",
		Success: true,
		Data:    nil,
	})
}
