package main

import (
	"fmt"
	"github.com/captainstorm21/web-app/pkg/handlers"
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
