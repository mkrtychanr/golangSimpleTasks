package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *Counter) Value() int {
	return c.counter
}

func main() {
	var wg sync.WaitGroup
	var c Counter
	ch := make(chan int, 10)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 2; i++ {
			go func(i int) {
				wg.Add(1)
				defer wg.Done()
				for data := range ch {
					fmt.Printf("[Worker %4d]: Data: %4d\n", i, data)
					c.Inc()
					time.Sleep(100 * time.Millisecond)
				}
			}(i)
		}
	}()

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println(c.Value())
}
