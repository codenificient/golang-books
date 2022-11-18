package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 19, 8}
	fmt.Println("Ranges in Go")
	fmt.Println("Indices and values:")
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// only element
	fmt.Println("Values only:")
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
