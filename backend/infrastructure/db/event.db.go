package db

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/luizemm/data-collector/domain/event"
	"github.com/luizemm/data-collector/infrastructure/log"
)

type EventDb struct {
	db *sql.DB
}

func NewEventDb(db *sql.DB) *EventDb {
	return &EventDb{db}
}

func (e *EventDb) Save(event event.EventInterface) error {
	query := "INSERT INTO event(event_type, data, timestamp) VALUES ($1, $2, $3) RETURNING id"

	parameters := []any{
		event.GetEventType(),
		event.GetData(),
		event.GetTimestamp().Format(time.UnixDate),
	}

	log.Logger.Info("Query: " + query, slog.Any("Parameters", parameters))

	var id int64

	err := e.db.QueryRow(query, parameters...).Scan(&id)

	if err != nil {
		return err
	}

	event.DefineId(uint64(id))

	return nil
}