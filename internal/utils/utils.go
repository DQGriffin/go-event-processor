package utils

import "github.com/DQGriffin/go-event-processor/internal/types"

func IsActionableEvent(event types.TelephonyEvent) bool {
	return (event.Body.Parties[0].Direction == "Outbound" && event.Body.Parties[0].Status.Code == "Setup") || (event.Body.Parties[0].Direction == "Inbound" && event.Body.Parties[0].Status.Code == "Answered")
}

func ExtractPayload(event types.TelephonyEvent) types.RecordingHandlerPayload {
	payload := types.RecordingHandlerPayload{
		UuiString: event.Body.TelephonySessionId,
		Direction: event.Body.Parties[0].Direction,
		From:      event.Body.Parties[0].From.PhoneNumber,
		To:        event.Body.Parties[0].To.PhoneNumber,
		AccountId: event.Body.AccountId,
	}
	return payload
}
