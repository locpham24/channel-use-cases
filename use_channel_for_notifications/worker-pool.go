package main

import (
	"fmt"
	"math/rand"
	"time"
)

func handleTask(i int, chanJob <-chan int, chanDone chan<- struct{}) {
	for {
		job, ok := <-chanJob
		if !ok {
			return
		}
		fmt.Println("worker", i, "started  job", job)
		latency := rand.Intn(3) + 2
		time.Sleep(time.Duration(latency) * time.Second)
		chanDone <- struct{}{}
		fmt.Printf("worker #%d done job #%d in %ds\n", i, job, latency)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 2, 3, 4, 5, 6}

	chanJob := make(chan int)
	chanDone := make(chan struct{}, 2) // at least 2 (numJob - numWorkers)

	numWorkers := 4
	fmt.Println("create worker...")
	for i := 1; i <= numWorkers; i++ {
		go handleTask(i, chanJob, chanDone)
	}

	fmt.Println("sending job...")
	for i := 0; i < len(arr); i++ {
		chanJob <- arr[i]
		fmt.Println("sent job", arr[i])
	}
	fmt.Println("close channel job...")
	close(chanJob)
	for a := 0; a < len(arr); a++ {
		<-chanDone
	}
}
