package storage

import (
	"testing"
	"os"
	"log"
)

func TestMain(s *testing.M) {
	err := Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(s.Run())
}

