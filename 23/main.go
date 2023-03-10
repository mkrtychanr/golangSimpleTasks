package main

import "fmt"

func deleteFromSlice(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = deleteFromSlice(s, 5)
	fmt.Println(s)
}
