package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RespondToWebhook(c *fiber.Ctx) error {
	var validationToken = c.Get("validation-token")

	if validationToken != "" {
		c.Set("validation-token", validationToken)
	}
	c.Set("content-type", "application/json")
	c.SendStatus(200)

	return c.Next()
}
