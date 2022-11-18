package main

import "fmt"

func main() {
	fmt.Println("Slices in Golang")
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s1, 1))
}

// Need to preserve the order of elements
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) // copy(dest, src)
	return slice[:len(slice)-1]
}

// order is not important
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
