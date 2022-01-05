package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName string
}

func main () {

	myMap := make(map[string]string)

	myMap["dog"] ="Sam"
	log.Println(myMap["dog"])

	myMap["cat"] ="Alisa"
	log.Println(myMap["cat"])

	myMap["dog"] ="Morgan"
	log.Println(myMap["dog"])


	myNameMap := make(map[string]User)

	me := User {
		FirstName: "Valentine",
		LastName: "Schwartz",
	}

	myNameMap["me"] = me 
	log.Println(myNameMap["me"].FirstName)

	// var  myNewVar float32
	// myNewVar = 11.2

	
}