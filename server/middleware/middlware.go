package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ApiMiddleware(c *fiber.Ctx) error {
	fmt.Println(c.IP())
	return c.Next()
}
