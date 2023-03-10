package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) string {
	data := []rune(str)
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return string(data)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	data := strings.Fields(str)
	var b strings.Builder
	b.WriteString(data[len(data)-1])
	for i := 1; i < len(data); i++ {
		real := len(data) - 1 - i
		b.WriteString(" ")
		b.WriteString(data[real])
	}
	str = b.String()
	fmt.Println(str)
}
