package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	// Create a router mux
	mux := chi.NewRouter()

	// Add middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// register routes
	mux.Get("/", app.Home)
	mux.Get("/about", app.About)
	mux.Get("/demomovies", app.AllDemoMovies)
	mux.Get("/movies", app.AllMovies)

	return mux
}
