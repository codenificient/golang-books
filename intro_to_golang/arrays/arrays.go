package main

import "fmt"

func main() {
	var arr [5]int
	arr[4] = 201
	fmt.Println(arr)
	arr[0] = 98
	arr[1] = 91
	arr[2] = 77
	arr[3] = 83
	fmt.Println(arr)

	var aray [5]float64
	aray[0] = 210
	aray[1] = 502
	aray[2] = 98.1
	aray[3] = 77.9
	aray[4] = 83.4
	fmt.Println(aray)

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += aray[i]
	}
	fmt.Println("The average of array is ", total/5)

	var totals float64 = 0
	for i := 0; i < len(aray); i++ {
		totals += aray[i]
	}
	fmt.Println("The average of array is", totals/float64(len(aray)))
}
