package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/main.db")
	if err != nil {
		return nil, err
	}

	return db, err
}
