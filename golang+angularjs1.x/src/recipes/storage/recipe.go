package storage

import (
	"recipes/models"

	goqu "gopkg.in/doug-martin/goqu.v4"
	_ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
)

func ListRecipes() ([]models.Recipe, error) {
	db := goqu.New("postgres", pgDB)
	var recipes []models.Recipe
	err := db.From("recipes").ScanStructs(&recipes)
	return recipes, err
}

func GetRecipe(id int) (models.Recipe, error) {
	db := goqu.New("postgres", pgDB)
	var recipe models.Recipe
	found, err := db.From("recipes").Where(goqu.I("id").Eq(id)).ScanStruct(&recipe)
	if err != nil {
		return recipe, err
	}
	if !found {
		return recipe, NotFoundError
	}
	return recipe, nil
}

func CreateRecipe(recipe models.Recipe) (int, error) {
	var id int
	db := goqu.New("postgres", pgDB)
	insert := db.From("recipes").Returning(goqu.I("id")).Insert(recipe)
	_, err := insert.ScanVal(&id)

	return id, err
}

func UpdateRecipe(id int, recipe models.Recipe) error {
	db := goqu.New("postgres", pgDB)
	_, err := db.From("recipes").Where(goqu.I("id").Eq(id)).Returning(goqu.I("id")).Update(recipe).Exec()

	if err != nil {
		return err
	}
	return nil
}

func DeleteRecipe(id int) error {
	db := goqu.New("postgres", pgDB)
	_, err := db.From("recipes").Where(goqu.I("id").Eq(id)).Delete().Exec()

	if err != nil {
		return err
	}
	return nil
}
