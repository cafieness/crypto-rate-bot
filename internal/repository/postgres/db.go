package postgres

import (
	"database/sql"
	"fmt"
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

	fmt.Println("Database connected")

	return db, nil
}
