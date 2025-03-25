package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/DQGriffin/go-event-processor/internal/types"
)

func SendToRecordingHandler(payload types.RecordingHandlerPayload) {
	log.Println("Sending payload to recording handler")

	json, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON body")
		log.Println(err.Error())
		return
	}

	res, err := http.Post(
		os.Getenv("RECORDING_HANDLER_HOST")+"/service/health",
		"application/json",
		bytes.NewBuffer(json),
	)

	if err != nil {
		log.Println("POST request failed")
		log.Println(err.Error())
		return
	}
	defer res.Body.Close()

	log.Println("POST Status: ", res.Status)

}
