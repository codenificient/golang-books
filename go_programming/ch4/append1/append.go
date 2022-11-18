package main

import "fmt"

func main() {
	fmt.Println("Appending Integer")
	var x, y []int
	for i := 0; i < 21; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d    cap=%d    %v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// grow by doubling for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // buitin function
	}
	// append new item
	z[len(x)] = y
	return z
}
