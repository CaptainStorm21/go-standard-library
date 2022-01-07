package handlers

import (	
	"github.com/captainstorm21/web-app/pkg/render"
	"net/http"
)

// Home is the about page handlers
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderPackage(w, "home.page.html")
}

// About is the about page handlers
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
