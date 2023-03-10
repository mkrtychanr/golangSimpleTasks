package main

import "fmt"

func binarySearch(data []int, target int) (int, bool) {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2

		if target == data[mid] {
			return mid, true
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	i, ok := binarySearch(data, 20)
	if !ok {
		fmt.Println("Sorry")
	} else {
		fmt.Println(i)
	}
}
