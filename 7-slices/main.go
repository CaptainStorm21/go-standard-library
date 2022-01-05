package main

import (
	"log"
	"sort"
)

func main() {

	var mySlice []string

	mySlice = append(mySlice, "Travis")
	mySlice = append(mySlice, "Mary")
	mySlice = append(mySlice, "Alice")

	log.Println(mySlice)

	sort.Strings(mySlice)
	log.Println(mySlice)

	sliceNumbers := []int{10, 11, 23, 43, 55, 990, 23, 11}
	log.Println(sliceNumbers)

	log.Println(sliceNumbers[3:6])

	names := []string{"Mary", "Jay", "Samantha", "Odessay"}
	log.Println(names)

	anyData := []interface{}{"foo", 10, true}
	log.Println(anyData)

}
