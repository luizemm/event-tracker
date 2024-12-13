package main

import (
	"github.com/luizemm/data-collector/infrastructure/db"
	"github.com/luizemm/data-collector/infrastructure/websocket"
)

func main() {
	db := db.OpenDatabase()
	defer db.Close()
	websocket.Init(db)
}