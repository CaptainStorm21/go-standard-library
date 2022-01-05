package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

var specialTitle = "CEO"
var SpecialSkill = "Web"

func main() {

	user := User{
		FirstName:   "Trevor",
		LastName:    "Williams",
		PhoneNumber: " 1 45 5534",
	}

	log.Println(user.FirstName, user.LastName)
	log.Println("Birthdate :", user.Birthdate)
	log.Println("Phone Number :", user.PhoneNumber)

}
