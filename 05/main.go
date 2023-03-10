package main

import (
	"fmt"
	"time"
)

func delta(startTime *time.Time) int {
	return int(time.Since(*startTime).Seconds())
}

func main() {
	n := 0
	fmt.Printf("Enter time: ")
	fmt.Scan(&n)
	startTime := time.Now()
	ch := make(chan int)

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	for i := 0; delta(&startTime) < n; i++ {
		ch <- i
	}
	close(ch)
}
