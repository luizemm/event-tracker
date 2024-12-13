package event

import "time"

type EventInterface interface {
	DefineId(id uint64) error
	GetId() uint64
	GetEventType() string
	GetData() string
	GetTimestamp() time.Time
}