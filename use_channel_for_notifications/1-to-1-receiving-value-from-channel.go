package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	done := make(chan struct{})
	values := make([]byte, 32*1024*1024)
	go func() {
		if _, err := rand.Read(values); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sort.Slice(values, func(i, j int) bool {
			if values[i] < values[j] {
				return true
			}
			return false
		})
		<-done
	}()

	done <- struct{}{}
	fmt.Printf("low:high %d %d \n", values[0], values[len(values)-1])
}
