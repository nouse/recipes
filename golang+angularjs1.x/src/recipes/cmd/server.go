package main

import (
	"net/http"
	"os"
	"recipes/handlers"
	"recipes/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	err := storage.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/recipes", handlers.RecipesRouter())

	http.ListenAndServe(":8080", r)
}
