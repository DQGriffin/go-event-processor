package routes

import (
	"github.com/DQGriffin/go-event-processor/internal/controllers"
	"github.com/DQGriffin/go-event-processor/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterWebhookRoutes(app *fiber.App) {
	system := app.Group("/api")

	system.Post("/webhook", middleware.RespondToWebhook, controllers.HandleWebhookRequest)
}
