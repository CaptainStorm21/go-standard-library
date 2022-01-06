package main

import (
	"log"
	"github.com/tsawler/myniceprogram/helpers"
)

const numPool = 1000

func CalcValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalcValue(intChan)

	num := <-intChan 
	log.Println(num)

}
