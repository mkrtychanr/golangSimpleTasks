package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch1 {
			ch2 <- num * num
		}
		close(ch2)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch2 {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}()

	data := []int{1, 2, 3, 4, 5, 6}

	for _, value := range data {
		ch1 <- value
	}
	close(ch1)
	wg.Wait()
}
