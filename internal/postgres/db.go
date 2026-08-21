package postgres

import (
	"database/sql"
	"log/slog"
)

func Connect(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("database connected")

	return db, nil
}
