package handlers

import (
	"github.com/CaptainStorm21/go-standard-library/tree/2-building-web-app/2-Building-web-app/pkg/render"
	"net/http"
)

// Home is the about page handlers
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handlers
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
