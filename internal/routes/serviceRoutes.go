package routes

import (
	"github.com/DQGriffin/go-event-processor/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterServiceRoutes(app *fiber.App) {
	system := app.Group("/service")

	system.Get("/health", controllers.HandleHealthRequest)
}
