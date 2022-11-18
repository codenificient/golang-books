package main

import "fmt"

func sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arrays[j] > arrays[j+1] { // swap
				var flag = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag
			}
		}
	}
}

func main() {
	var scores = []int{50, 60, 83, 95, 70}
	var length = len(scores)
	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
