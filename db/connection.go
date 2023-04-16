package db

import (
	"os"
	"database/sql"
)

var Connection *sql.DB

func Connect() (err error) {
	Connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	return nil
}