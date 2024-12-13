package db_test

import (
	"database/sql"
	"log"
	"testing"
	"time"

	ev "github.com/luizemm/data-collector/domain/event"
	"github.com/luizemm/data-collector/infrastructure/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE event (
		"id" integer primary key,
		"event_type" string,
		"data" string,
		"timestamp" string
	);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestEventDbSave(t *testing.T) {
	setUp()
	defer Db.Close()
	
	eventDb := db.NewEventDb(Db)

	event := ev.NewEvent(ev.EventProps{
		EventType: "eventType",
		Data: "Data",
		Timestamp: time.Now(),
	})

	err := eventDb.Save(event)

	require.Nil(t, err)

	var rowResult struct{
		Id uint64
		EventType string
		Data string
		Timestamp string
	}

	err = Db.QueryRow(
		"SELECT * FROM event WHERE id = ?", event.GetId(),
	).Scan(
		&rowResult.Id, &rowResult.EventType,
		&rowResult.Data, &rowResult.Timestamp,
	)

	timestamp, err := time.Parse(time.UnixDate, rowResult.Timestamp)

	require.Nil(t, err)

	eventResult := ev.NewEvent(ev.EventProps{
		Id: rowResult.Id,
		EventType: rowResult.EventType,
		Data: rowResult.Data,
		Timestamp: timestamp,
	})

	require.Nil(t, err)
	require.NotNil(t, event.GetId(), eventResult.GetId())
	require.Equal(t, event.GetEventType(), eventResult.GetEventType())
	require.Equal(t, event.GetData(), eventResult.GetData())
	require.Equal(t, event.GetTimestamp().Format(time.UnixDate), eventResult.GetTimestamp().Format(time.UnixDate))
}