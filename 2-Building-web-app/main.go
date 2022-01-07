package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the about page handlers
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a homepage")
}

// About is the about page handlers
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 3)
	fmt.Fprintf(w, "this is About page")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page %d", sum))
}

// Adding values
func addValue(x, y int) int {
	return x + y
}

// Main is the application function
func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the server on %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
