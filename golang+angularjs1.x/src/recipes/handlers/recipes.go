package handlers

import (
	"context"
	"net/http"
	"recipes/models"
	"recipes/storage"

	"encoding/json"
	"io/ioutil"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"strconv"
)

func RecipesRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", ListRecipes)
	r.Post("/", CreateRecipe)

	r.Route("/{recipeID}", func(r chi.Router) {
		r.Use(RecipeCtx)            // Load the *Recipe on the request context
		r.Get("/", GetRecipe)       // GET /recipes/123
		r.Put("/", UpdateRecipe)    // GET /recipes/123
		r.Delete("/", DeleteRecipe) // GET /recipes/123
	})
	return r
}

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := storage.ListRecipes()
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}
	if len(recipes) == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}
	render.JSON(w, r, recipes)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := r.Context().Value("recipe").(*models.Recipe)
	render.JSON(w, r, recipe)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newRecipe := models.Recipe{}

	requestBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(requestBody, &newRecipe)
	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}

	id, err := storage.CreateRecipe(newRecipe)

	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}
	newRecipe.ID = id

	render.JSON(w, r, newRecipe)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	recipe := r.Context().Value("recipe").(*models.Recipe)

	newRecipe := models.Recipe{}

	requestBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(requestBody, &newRecipe)
	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}

	err = storage.UpdateRecipe(recipe.ID, newRecipe)

	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}

	render.JSON(w, r, newRecipe)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := r.Context().Value("recipe").(*models.Recipe)
	err := storage.DeleteRecipe(recipe.ID)

	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}

	render.JSON(w, r, map[string]int{"ID": recipe.ID})
}

func RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var recipe models.Recipe
		var err error

		if recipeID := chi.URLParam(r, "recipeID"); recipeID != "" {
			id, err := strconv.Atoi(recipeID)
			if err != nil {
				render.Status(r, 400)
				render.PlainText(w, r, err.Error())
				return
			}
			if id <= 0 {
				render.Status(r, 400)
				render.PlainText(w, r, "ID should be positive integer")
				return
			}
			recipe, err = storage.GetRecipe(id)
		} else {
			render.Status(r, 400)
			render.PlainText(w, r, "need to specify recipeID")
			return
		}
		if err != nil {
			render.Status(r, 404)
			render.PlainText(w, r, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "recipe", &recipe)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
