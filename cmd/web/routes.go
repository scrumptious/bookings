package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/scrumptious/bookings/pkg/config"
	"github.com/scrumptious/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(app.Session.LoadAndSave)
	r.Use(CsrfMid)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	return r
}
