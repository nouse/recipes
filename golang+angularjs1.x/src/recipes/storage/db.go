package storage

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

var (
	pgDB          *sql.DB
	NotFoundError = errors.New("record not found")
)

func Connect(connStr string) error {
	var err error
	if pgDB != nil {
		err = pgDB.Close()
	}
	if err != nil {
		return err
	}
	pgDB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}

