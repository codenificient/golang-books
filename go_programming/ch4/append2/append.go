package main

import "fmt"

func main() {
	fmt.Println("Append several Values")

}

func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// expend z to at least zlen
	copy(z[len(x):], y)
	return z
}
