package main

import (
	"net/http"
)

// Home is the about page handlers
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the about page handlers
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
