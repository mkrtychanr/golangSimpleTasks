package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]uint8)

	for _, value := range data {
		set[value] = 0
	}

	for key := range set {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
}
