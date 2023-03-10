package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// первый способ – просто закрыть канал, который использует горутина и соответствующим способом это поймать
	ch := make(chan int)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("[Goroutine 1]: Done")
				return
			}
		}
	}()

	ch <- 1
	ch <- 1
	ch <- 1
	close(ch)
	wg.Wait()

	// второй способ - создать специальный канал, который будет жать сигнала остановки
	quit := make(chan bool)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("[Goroutine 2]: Done")
				return
			default:
				// код с lipsum.com
			}
		}
	}()

	quit <- true
	wg.Wait()

}
