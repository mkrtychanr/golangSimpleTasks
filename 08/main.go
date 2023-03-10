package main

import "fmt"

type bitset int64

func (s *bitset) setBit(pos uint8, bit uint8) {
	if pos > 64 {
		return
	}
	var ans int64
	ans = int64(*s)
	if bit == 0 {
		var mask int64
		mask = ^(1 << pos)
		ans &= mask
	} else {
		ans |= (1 << pos)
	}
	*s = bitset(ans)
}

func main() {
	var a bitset
	a = 0
	a.setBit(0, 1)
	fmt.Println(a)
	a.setBit(1, 1)
	fmt.Println(a)
	a.setBit(0, 0)
	a.setBit(1, 0)
	fmt.Println(a)
}
