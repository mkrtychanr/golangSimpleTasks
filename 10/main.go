package main

import (
	"fmt"
	"math"
)

func getKey(value float64) int64 {
	return int64(value - math.Mod(value, 10))
}

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mp := make(map[int64][]float64)

	for _, value := range data {
		key := getKey(value)
		_, ok := mp[key]
		if !ok {
			mp[key] = make([]float64, 0)
		}
		mp[key] = append(mp[key], value)
	}

	for key, slice := range mp {
		fmt.Printf("%d:%v\n", key, slice)
	}
}
