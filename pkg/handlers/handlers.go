package handlers

import (
	"github.com/scrumptious/bookings/pkg/config"
	"github.com/scrumptious/bookings/pkg/models"
	"github.com/scrumptious/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (rp *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "ip", remoteIP)
	stringMap := map[string]string{
		"test": "Testing template data huhuhuh",
	}
	render.RenderTemplate(rw, "home.page.tmpl.html", &models.TemplateData{StringMap: stringMap})
}

// About is the about page handler
func (rp *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"ip": rp.App.Session.GetString(r.Context(), "ip"),
	}

	render.RenderTemplate(rw, "about.page.tmpl.html", &models.TemplateData{StringMap: stringMap})
}
