package main

import (
	"log"
	"time"
)

func worker(id int, ready, done chan T) {
	log.Print("Worker#", id, " waits.")
	<-ready
	log.Print("Worker#", id, " starts.")

	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")

	done <- T{}
}

type T struct{}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	time.Sleep(time.Second * 3 / 2)
	close(ready)
	<-done
	<-done
	<-done
}
