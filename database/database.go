package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(name, datasource string) error {
	var err error
	db, err = sql.Open(name, datasource)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	db.Close()
}
