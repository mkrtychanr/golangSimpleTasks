package main

import (
	"fmt"
	"sync"
)

type pair struct {
	First  int
	Second string
}

func main() {

	var (
		wg sync.WaitGroup
		mu sync.Mutex
		n  int
	)

	dataSlice := []pair{{1, "one"}, {2, "second"}, {3, "third"}, {4, "fourth"}, {5, "fifth"}, {6, "sixth"}}

	ch := make(chan pair)

	data := make(map[int]string)

	n = 2

	go func() {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for value := range ch {
					mu.Lock()
					data[value.First] = value.Second
					mu.Unlock()
				}
			}()
		}
	}()

	for _, value := range dataSlice {
		ch <- value
	}

	close(ch)
	wg.Wait()

	for key, value := range data {
		fmt.Printf("%v: %v\n", key, value)
	}

}
