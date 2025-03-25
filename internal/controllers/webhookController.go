package controllers

import (
	"log"

	"github.com/DQGriffin/go-event-processor/internal/services"
	"github.com/DQGriffin/go-event-processor/internal/types"
	"github.com/DQGriffin/go-event-processor/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleWebhookRequest(c *fiber.Ctx) error {
	var err error
	var event types.TelephonyEvent

	if err := c.BodyParser(&event); err != nil {
		log.Println("Bad JSON body received")
		log.Printf("Received event: %+v\n", event)
		return err
	}

	if utils.IsActionableEvent(event) {
		payload := utils.ExtractPayload(event)
		log.Println("Actionable event received", payload)

		if err := services.SendToRecordingHandler(payload); err != nil {
			log.Println("Failed to send payload to recording handler")
			log.Println(err.Error())
			return err
		}
	}

	return err
}
