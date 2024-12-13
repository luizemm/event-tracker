package db

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/luizemm/data-collector/infrastructure/env"
	"github.com/luizemm/data-collector/infrastructure/log"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open(os.Getenv(env.DB_DRIVER_NAME), os.Getenv(env.DB_URL_CONNECTION))

	if err != nil {
		log.Logger.Error("Open database", slog.Any("error", err))
		os.Exit(1)
	}

	return db
}