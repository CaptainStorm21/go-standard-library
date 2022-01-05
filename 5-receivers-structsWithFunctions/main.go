package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}



func main() {

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// log.Println("MyVar set to ", myVar.FirstName)
	// log.Println("MyVar set to ", myVar2.FirstName)

	log.Println("MyVar func set to ", myVar.printFirstName())
	log.Println("MyVar func set to ", myVar2.printFirstName())
}