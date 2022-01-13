package handlers

import (
	"net/http"
	"github.com/captainstorm21/go-web-app/pkg/render"
		"github.com/captainstorm21/go-web-app/pkg/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a, 
	}
}

func NewHandlers (r *Repository){
	Repo = r 
}


// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
