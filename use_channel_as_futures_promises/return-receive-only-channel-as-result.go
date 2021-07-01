package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)
// assume that need various latency to gen this number
func genNumber() <-chan int {
	chanNumber := make(chan int)

	go func(){
		time.Sleep(time.Second * time.Duration(3)) // example latency because of business rule, query database
		randNumber := int(rand.Int31n(100))
		log.Println("number:", randNumber)
		chanNumber <- randNumber
	}()
	return chanNumber
}

func sum(a, b int) int {
	return a + b
}
func main(){
	log.Println("START")
	rand.Seed(time.Now().UnixNano())

	a, b := genNumber(), genNumber()
	fmt.Println("SUM = ", sum(<-a, <-b))
	log.Println("END")
}