package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	n := 0
	fmt.Printf("Enter the number of workers: ")
	fmt.Scan(&n)
	ch := make(chan string, n)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for data := range ch {
				mu.Lock()
				fmt.Printf("[Worker %4d]: %s\n", i, data)
				mu.Unlock()
				time.Sleep(10 * time.Second)
			}
			fmt.Printf("[Worker %4d]: Done\n", i)
		}(i)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		close(ch)
		wg.Wait()
		os.Exit(0)
	}()
	for {
		data := ""
		fmt.Printf("Enter data: ")
		fmt.Scan(&data)
		ch <- data
	}
}
