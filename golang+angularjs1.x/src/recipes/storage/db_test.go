package storage

import (
	"testing"
	"os"
	"log"
	"github.com/DATA-DOG/go-txdb"
	"database/sql"
)

func TestMain(s *testing.M) {
	txdb.Register("txdb", "postgres", os.Getenv("DATABASE_URL"))
	var err error
	pgDB, err = sql.Open("txdb", "identifier")
	defer pgDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(s.Run())
}

