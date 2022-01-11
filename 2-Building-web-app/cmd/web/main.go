package main

import (
	"fmt"
	"net/http"
	"github.com/captainstorm21/go-web-app/pkg/handlers"
)

const portNumber = ":8080"

// Main is the application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the server on %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
