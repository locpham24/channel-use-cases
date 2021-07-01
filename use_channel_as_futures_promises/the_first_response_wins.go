package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func source(c chan<- int) {
	delay := rand.Int31n(10)
	number := int(rand.Int31n(100))
	fmt.Printf("delay: %d number:%d\n", delay,number)
	time.Sleep(time.Second * time.Duration(delay))
	c <- number
}

func main(){
	rand.Seed(time.Now().UnixNano())

	log.Println("START")
	startTime := time.Now()
	c := make(chan int, 5)
	for i:= 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println("result", rnd)
	time.Sleep(20 * time.Second)
	log.Println("END")
}