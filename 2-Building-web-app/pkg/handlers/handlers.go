package handlers

import (
<<<<<<< Updated upstream
	"net/http",
	"github.com/captainstorm21/"
)

// Home is the about page handlers
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderPackage(w, "home.page.html")
}

// About is the about page handlers
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
=======
	"github.com/CaptainStorm21/go-standard-library/tree/2-building-web-app/2-Building-web-app/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
>>>>>>> Stashed changes
}
