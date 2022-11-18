package main

import "fmt"

func main() {
	xs := []float64{98, 77, 82, 0192.2}
	fmt.Println("The average of these numbers", xs, "is", average(xs))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
