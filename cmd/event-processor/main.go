package main

import (
	"log"
	"os"

	"github.com/DQGriffin/go-event-processor/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load from environment file")
		log.Println(err.Error())
	}

	app := fiber.New()
	app.Use(recover.New())

	routes.RegisterServiceRoutes(app)
	routes.RegisterWebhookRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT variable is not set. Using default value of 3000")
		port = "3000"
	}

	addr := ":" + port

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}

}
