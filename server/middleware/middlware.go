package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// APIMiddleware is the layer between the request and the response
func APIMiddleware(c *fiber.Ctx) error {
	// usefulData := map[string]interface{}{
	// 	"Method":     c.Method(),
	// 	"IP":         c.IP(),
	// 	"Hostname":   c.Hostname(),
	// 	"RemoteAddr": c.Context().RemoteAddr(),
	// 	"RemoteIP":   c.Context().RemoteIP(),
	// 	"LocalAddr":  c.Context().LocalAddr(),
	// 	"LocalIP":    c.Context().LocalIP(),
	// 	"ClientIP":   realip.FromRequest(c.Context()),
	// }

	// for key, value := range usefulData {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	return c.Next()
}
