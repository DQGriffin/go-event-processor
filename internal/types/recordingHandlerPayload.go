package types

type RecordingHandlerPayload struct {
	UuiString string `json:"uuiString"`
	Direction string `json:"direction"`
	From      string `json:"from"`
	To        string `json:"to"`
	AccountId string `json:"accountId"`
}
