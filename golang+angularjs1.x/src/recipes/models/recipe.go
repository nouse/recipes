package models

type Recipe struct {
	ID           int    `db:"id" json:"id" goqu:"skipupdate,skipinsert"`
	Title        string `db:"title" json:"title"`
	Description  string `db:"description" json:"description"`
	Ingredients  jsonb  `db:"ingredients" json:"ingredients"`
	Instructions string  `db:"instructions" json:"instructions"`
}

