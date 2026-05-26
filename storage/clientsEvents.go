package storage

import (
	"time"
	"fmt"
)

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

type ConnectionEvent struct {
	ClientEvent
	topics []string
}

type Event interface{
	ToString()string
}

func NewConnectionEvent(clientId string, topics []string, isConnectionEvent bool, ts time.Time) *ConnectionEvent {
	var status EventType
	if isConnectionEvent {
		status = CONNECTION
	} else {
		status = DISCONNECTION
	}

	event := ConnectionEvent{
		ClientEvent: ClientEvent{
			ClientId: clientId,
			Timestamp: ts,
			Status: status,
		},
		topics: topics,
	}
	return &event
}

func (c *ConnectionEvent) ToString()string{
	return fmt.Sprintf("%v: %s %v - %v", c.Timestamp, c.ClientId, c.Status, c.topics)
}

func NewClientEvent(clientId string, status EventType, eventTime time.Time) *ClientEvent {
	client := ClientEvent{
		Timestamp: eventTime,
		Status: status,
		ClientId: clientId,
	}
	return &client
}
