package controllers

import (
	"errors"
	"log"

	"github.com/DQGriffin/go-event-processor/internal/services"
	"github.com/DQGriffin/go-event-processor/internal/types"
	"github.com/gofiber/fiber/v2"
)

func HandleWebhookRequest(c *fiber.Ctx) error {
	var err error
	var event types.TelephonyEvent

	if err := c.BodyParser(&event); err != nil {
		log.Println("Bad JSON body received")
		err = errors.New("bad JSON")
		return err
	}

	if (event.Body.Parties[0].Direction == "Outbound" && event.Body.Parties[0].Status.Code == "Setup") || (event.Body.Parties[0].Direction == "Inbound" && event.Body.Parties[0].Status.Code == "Answered") {
		payload := types.RecordingHandlerPayload{
			UuiString: event.Body.TelephonySessionId,
			Direction: event.Body.Parties[0].Direction,
			From:      event.Body.Parties[0].From.PhoneNumber,
			To:        event.Body.Parties[0].To.PhoneNumber,
			AccountId: event.Body.AccountId,
		}

		log.Println("Actionable event received", payload)
		services.SendToRecordingHandler(payload)
	}

	return err
}
