package main

import "fmt"

func main() {
	array := [6]float64{201, 32, 430.0, 21.29, 897.07}
	var total float64 = 0
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	fmt.Println("This array", array, "has an average of", total/float64(len(array)))

	x := array[0:3]
	fmt.Println(x)
	fmt.Println(array[2:])
	slice1 := []int{1, 2, 3, 4}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 is", slice1, "slice2 is", slice2)
}
