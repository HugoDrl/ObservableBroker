package storage

import "time"

type EventType int
const (
	CONNECTION EventType = iota
	DISCONNECTION
)
type ClientEvent struct {
	Timestamp time.Time
	Status EventType
	ClientId string
}

func NewClientEvent(clientId string, status EventType, eventTime time.Time) *ClientEvent {
	client := ClientEvent{
		Timestamp: eventTime,
		Status: status,
		ClientId: clientId,
	}
	return &client
}
