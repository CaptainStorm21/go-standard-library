package main

import (
	"net/http"

	"github.com/captainstorm21/go-web-app/pkg/config"
	"github.com/captainstorm21/go-web-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer())

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.Home)
	return mux
}
