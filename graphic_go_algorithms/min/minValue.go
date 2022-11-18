package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex = 0
	for i := 1; i < length; i++ {
		if arrays[minIndex] > arrays[i] { // swap
			minIndex = i
		}
	}
	return arrays[minIndex]
}

func main() {
	var scores = []int{50, 60, 83, 95, 70}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Println("Minimum value: ", minValue)
}
