package main

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	goqu "gopkg.in/doug-martin/goqu.v4"
	_ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
)

type JSONB json.RawMessage

func (j JSONB) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j JSONB) IsNull() bool {
	return len(j) == 0 || bytes.Equal(j, []byte("null"))
}

type Recipe struct {
	ID          int    `db:"id" goqu:"skipupdate"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Ingredients JSONB  `db:"ingredients"`
}

func main() {
	pgDb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	db := goqu.New("postgres", pgDb)
	logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
	db.Logger(logger)
	var recipe Recipe
	found, err := db.From("recipes").Where(goqu.I("id").Eq(1)).ScanStruct(&recipe)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if found {
		s, _ := json.Marshal(recipe)
		fmt.Printf(string(s))
	}

	// Update

	recipe.Title = "New Title"
	_, err = db.From("recipes").Where(goqu.I("id").Eq(1)).Returning(goqu.I("id")).Update(recipe).Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
