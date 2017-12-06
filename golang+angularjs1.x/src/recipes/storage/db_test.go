package storage

import (
	"testing"
	"os"
	"github.com/DATA-DOG/go-txdb"
	"database/sql"
)

func openTestDB() error {

	var err error
	if pgDB != nil {
		err = pgDB.Close()
	}
	if err != nil {
		return err
	}
	pgDB, err = sql.Open("txdb", "identifier")
	return err
}

func TestMain(s *testing.M) {
	txdb.Register("txdb", "postgres", os.Getenv("DATABASE_URL"))
	defer func() {
		if pgDB != nil {
			pgDB.Close()
		}
	}()
	os.Exit(s.Run())
}

