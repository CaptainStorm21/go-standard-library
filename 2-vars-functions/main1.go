package main

import (
	"log"
	"time"
)

var s2 string
var word = "word"

//interface

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

func main() {
	var s2 = "six"

	word := "eight"

	log.Println("this is s2", s2)
	log.Println("this is word", word)

	saySomething("xxx")
}

func saySomething(s string) (string, string) {
	log.Println("s from saySomething", s)
	return word, "world"
}
