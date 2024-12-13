package event

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type EventDto struct {
	EventType string `json:"event_type"`
	Data string `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *EventDto) Json() string{
	dtoJson, _ := json.Marshal(e)

	return string(dtoJson)
}

func (e *EventDto) Validate() error{
	var err error = nil

	if strings.TrimSpace(e.EventType) == "" {
		err = errors.Join(err, errors.New("event_type is required"))
	}

	if e.Data == "" {
		err = errors.Join(err, errors.New("data is required"))
	}

	if(!json.Valid([]byte(e.Data))) {
		err = errors.Join(err, errors.New("data must be in JSON format"))
	}
		
	return err
}