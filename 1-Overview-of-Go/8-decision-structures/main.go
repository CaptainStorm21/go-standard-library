package main

import (
	"log"
)

func main() {
	var isTrue bool
	
	isTrue = true
	
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if (cat == "cat"){
		log.Println("the cat is cat")
	} else {
		log.Println("the cat must be a dog")
	}
	

	myNum := 200
	isBoolTrue := false 

	if (myNum > 100 && !isBoolTrue){
		log.Println("my num is greater than 100 ")
	} else if myNum == 200 {
		log.Println("Number is 200")
	} else {
		log.Println("Where is the number")
	}


	myVar := "dolphin"
	switch myVar {

	case "cat": 
		log.Println("cat is the cat")

	case "dog":
		log.Println("German Shepherd")

	default: log.Println("check the wikipedia")
	}

}
