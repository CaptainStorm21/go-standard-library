package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var what2Say string
	what2Say = "Cruel Cat"

	fmt.Println("My name is " + what2Say + "!")

	var i int 
	i = 90
	fmt.Println(i, " is my number");	
	
	whatWasSaid :=saySomething()
	fmt.Println("The function returned, " + whatWasSaid)

	helloToMars, helloToVenus := greetingsToNeighbors()
	fmt.Println("The greetings were sent to " + helloToMars + " & " + helloToVenus)
}



func saySomething() string {
	return "something was said!"
}

func greetingsToNeighbors() (string, string){
	return  "Mars", "Venus"
}

