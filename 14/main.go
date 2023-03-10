package main

import (
	"fmt"
	"strings"
)

func typeDetection(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	default:
		t := fmt.Sprintf("%T", v)
		if strings.Compare(t[0:4], "chan") == 0 {
			fmt.Println("Channel")
		} else {
			fmt.Println("...")
		}
	}
}

func main() {
	var a int64
	typeDetection(a)
}
