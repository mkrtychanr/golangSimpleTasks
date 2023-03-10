package main

import "fmt"

func main() {
	set1 := make(map[int]uint8)
	set2 := make(map[int]uint8)

	for i := 0; i < 25; i++ {
		set1[i] = 0
	}

	for i := 15; i < 30; i++ {
		set2[i] = 0
	}

	intersectionSet := make(map[int]bool)

	for key := range set1 {
		intersectionSet[key] = false
	}

	for key := range set2 {
		if _, ok := intersectionSet[key]; ok {
			intersectionSet[key] = true
		}
	}

	intersection := make([]int, 0)

	for key, value := range intersectionSet {
		if value {
			intersection = append(intersection, key)
		}
	}

	for _, value := range intersection {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

}
