package main

import (
	"github.com/CaptainStorm21/go-standard-library/tree/2-building-web-app/2-Building-web-app/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Main is the application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the server on %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
