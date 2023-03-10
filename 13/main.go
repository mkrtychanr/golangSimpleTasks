package main

import "fmt"

func main() {
	var a, b int64

	a = 64
	b = 128

	a, b = b, a

	fmt.Printf("a: %d\nb: %d\n", a, b)
}
