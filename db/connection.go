package db

import "database/sql"

var Connection *sql.DB

func Connect() (err error) {
	Connection, err = sql.Open("postgres", "postgres://sdil:h27SCilrpdjK@ep-throbbing-surf-206751-pooler.ap-southeast-1.aws.neon.tech/mukmin-life")
	if err != nil {
		return err
	}

	return nil
}