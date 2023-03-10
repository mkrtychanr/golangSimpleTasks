package main

import "fmt"

func reverse(str string) string {
	data := []rune(str)
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return string(data)
}

func main() {
	var data string
	fmt.Scan(&data)
	data = reverse(data)
	fmt.Println(data)
}
