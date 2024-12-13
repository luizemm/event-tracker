package event

import (
	"errors"
	"time"
)

type Event struct {
	id uint64
	eventType string
	data string
	timestamp time.Time
}

type EventProps struct {
	Id uint64
	EventType string
	Data string
	Timestamp time.Time
}

func NewEvent(event EventProps) *Event {
	return &Event{
		id: event.Id,
		eventType: event.EventType,
		data: event.Data,
		timestamp: event.Timestamp.UTC(),
	}
}

func (e *Event) DefineId(id uint64) error {
	if e.id != 0 {
		return errors.New("Id is already defined")
	}

	e.id = id

	return nil
}

func (e Event) GetId() uint64 {
	return e.id
}

func (e Event) GetEventType() string {
	return e.eventType
}

func (e Event) GetData() string {
	return e.data
}

func (e Event) GetTimestamp() time.Time {
	return e.timestamp
}