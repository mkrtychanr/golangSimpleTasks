package main

func createHugeString(size uint64) string {
	return string(make([]rune, size))
}

func getSlice(v string, from, to uint64) string {
	data := []rune(v)
	slice := make([]rune, 0)
	for i := from; i < to; i++ {
		slice = append(slice, data[i])
	}
	return string(slice)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = getSlice(v, 0, 100)
}

func main() {
	someFunc()
}
