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

	// register routes
	mux.Get("/", app.Hello)
	mux.Get("/about", app.About)

	return mux
}
