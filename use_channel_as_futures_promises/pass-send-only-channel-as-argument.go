package main

import(
	"fmt"
	"log"
	"math/rand"
	"time"
)

func genNumberByChannelArgument(chanNumber chan<- int) {
	time.Sleep(time.Second * 3)
	num := int(rand.Int31n(100))
	fmt.Println("number:", num)
	chanNumber <- num
}

func calSum(a, b int) int {
	return a + b
}

func main(){
	log.Println("START")
	rand.Seed(time.Now().UnixNano())

	chanA := make(chan int)
	chanB := make(chan int)

	go genNumberByChannelArgument(chanA)
	go genNumberByChannelArgument(chanB)

	fmt.Println("SUM: " , calSum(<-chanA, <-chanB))


	/*chanNumber := make(chan int, 2)

	go genNumberByChannelArgument(chanNumber)
	go genNumberByChannelArgument(chanNumber)

	fmt.Println("SUM: " , calSum(<-chanNumber, <-chanNumber))*/

	log.Println("END")
}
