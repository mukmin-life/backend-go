package db

import (
	"os"
	"time"
	"database/sql"
)

var Connection *sql.DB

func Connect() (err error) {
	Connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	Connection.SetMaxOpenConns(200)
	Connection.SetMaxIdleConns(200)
	Connection.SetConnMaxLifetime(5*time.Minute)


	return nil
}