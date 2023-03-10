package main

import (
	"fmt"

	"github.com/mkrtychanr/wildberries_L1/24/point"
)

func main() {
	var (
		x float64
		y float64
	)

	fmt.Println("1st point")
	fmt.Printf("Enter x: ")
	fmt.Scan(&x)
	fmt.Printf("Enter y: ")
	fmt.Scan(&y)
	first := point.NewPoint(x, y)

	fmt.Println("2nd point")
	fmt.Printf("Enter x: ")
	fmt.Scan(&x)
	fmt.Printf("Enter y: ")
	fmt.Scan(&y)

	second := point.NewPoint(x, y)

	fmt.Println(first.GetDistance(&second))
}
