package controllers

import "github.com/gofiber/fiber/v2"

func HandleHealthRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "OK",
	})
}
