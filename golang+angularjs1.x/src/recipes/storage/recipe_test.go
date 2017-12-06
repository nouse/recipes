package storage

import (
	"testing"
	"recipes/models"
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func TestCreateRecipe(t *testing.T) {
	err := openTestDB()
	if err != nil {
		t.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	recipe := models.Recipe{
		ID: 1,
		Title: fmt.Sprintf("T%d", rand.Int()),
		Ingredients: []byte(`[{"amount":10},{"amount":20}]`),
	}

	id, err := CreateRecipe(recipe)
	if err != nil {
		t.Errorf("Fail to create recipe, error:%s", err)
	}
	if id <=0 {
		t.Errorf("Unexpected id, actual:%d", id)
	}
}

func TestGetRecipe(t *testing.T) {
	err := openTestDB()
	if err != nil {
		t.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	recipe := models.Recipe{
		ID: 1,
		Title: fmt.Sprintf("T%d", rand.Int()),
		Ingredients: []byte(`[{"amount": 10}, {"amount": 20}]`),
	}

	id, err := CreateRecipe(recipe)
	if err != nil {
		t.Errorf("Fail to create recipe, error:%s", err)
	}
	if id <=0 {
		t.Errorf("Unexpected id, actual:%d", id)
	}

	newRecipe, err := GetRecipe(id)
	if err != nil {
		t.Errorf("Fail to get recipe, error:%s", err)
	}
	if !bytes.Equal(newRecipe.Ingredients, recipe.Ingredients) {
		t.Errorf("Unexpected ingredients, actual: %s", string(newRecipe.Ingredients))
	}
}
