package main

import (
	"github.com/luizemm/event-tracker/infrastructure/db"
	"github.com/luizemm/event-tracker/infrastructure/websocket"
)

func main() {
	db := db.OpenDatabase()
	defer db.Close()
	websocket.Init(db)
}