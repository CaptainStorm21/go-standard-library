package main

import (
	"errors"
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

// Divide
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0)
	if err != nil {
		fmt.Fprint(w, "Cannog divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided %f is %f", 900.0, 0.0, f))
}

func divideValue(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
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
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the server on %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
