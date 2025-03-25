package types

import (
	"encoding/json"
	"fmt"
)

type TelephonyEvent struct {
	Uuid           string             `json:"uuid"`
	SubscriptionId string             `json:"subscriptionId"`
	OwnerId        string             `json:"ownerId"`
	Event          string             `json:"event"`
	Timestamp      string             `json:"timestamp"`
	Body           TelephonyEventBody `json:"body"`
}

type TelephonyEventBody struct {
	Sequence           uint8            `json:"sequence"`
	SessionId          string           `json:"sessionId"`
	TelephonySessionId string           `json:"telephonySessionId"`
	ServerId           string           `json:"serverId"`
	EventTime          string           `json:"eventTime"`
	AccountId          string           `json:"accountId"`
	Parties            []TelephonyParty `json:"parties"`
}

type TelephonyParty struct {
	ExtensionId string               `json:"extensionId"`
	Id          string               `json:"id"`
	Direction   string               `json:"direction"`
	From        TelephonyParticipant `json:"from"`
	To          TelephonyParticipant `json:"to"`
	Status      TelephonyStatus      `json:"status"`
}

type TelephonyParticipant struct {
	PhoneNumber string  `json:"phoneNumber"`
	Name        string  `json:"name"`
	ExtensionId *string `json:"extensionId,omitempty"`
}

type TelephonyStatus struct {
	Code string `json:"code"`
}

func (e TelephonyEvent) String() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Sprintf("TelephonyEvent<error: %v>", err)
	}
	return string(b)
}
