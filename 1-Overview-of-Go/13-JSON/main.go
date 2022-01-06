package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog string `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Jason",
			"last_name": "Williams",
			"hair_color": "Light brown",
			"has_dog": true
		},
		{
			"first_name": "Jaycee",
			"last_name": "Willis",
			"hair_color": "Blonde",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling")
	}
	log.Printf( "unmarshalled: %v", unmarshalled)


}