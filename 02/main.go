package main

import (
	"fmt"
	"sync"
	"time"
)

type Councurrency struct {
	mu   sync.Mutex
	data []int
}

func NewCouncurrency(data []int) *Councurrency {
	c := Councurrency{
		data: make([]int, len(data)),
	}
	for i, value := range data {
		c.data[i] = value
	}
	return &c
}

func (c *Councurrency) GetSquare(wg *sync.WaitGroup, index int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Start %d\n", index)
	c.data[index] = c.data[index] * c.data[index]
	time.Sleep(1 * time.Second)
	fmt.Printf("End %d\n", index)
	wg.Done()
}

func main() {
	sl := []int{2, 4, 6, 8, 10}
	c := NewCouncurrency(sl)
	var wg sync.WaitGroup
	for i := range sl {
		wg.Add(1)
		go c.GetSquare(&wg, i)
	}
	wg.Wait()
	for _, value := range c.data {
		fmt.Printf("%v ", value)
	}
	fmt.Println()
}
