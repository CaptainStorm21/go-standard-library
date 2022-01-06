package main

import (
	"log"
)

func main (){
	// for i :=0 ; i <= 5; i++{
	// 	log.Println(i)
	// }

	animals := []string {
		"dog",
		"cat",
		"elephant",
		"zebra",
	}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}


	zooAnimals := make(map[string]string)

	zooAnimals["zebra"] = "Milly"
	zooAnimals["camel"] = "Cairo"

	for _, zoo := range zooAnimals {
		log.Println(zoo)
	}


	var firstLine = "Once upon a time there was a moon"
	// firstLine = "x"

	for i, l := range firstLine {
		log.Println( i, ":", l)
	}


	type User struct {
		FirstName string
		LastName string
		Age int
		email string
		Graduate bool
	}

	var users []User
	users = append(users, User{"John", "Smith", 78, "whatever@yahoo.com", false})
	users = append(users, User{"Alice", "Smith", 78, "whatever@yahoo.com", true})
	users = append(users, User{"Naty", "Smith", 78, "whatever@yahoo.com", true})
	users = append(users, User{"Victoria", "Smith", 78, "whatever@yahoo.com", false})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Graduate)
	}

}