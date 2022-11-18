package main

import "fmt"

func main() {
	fmt.Println("Deep equality")
	a := []string{"banana", "apple", "mango", "orange", "pear", "papaya"}
	b := []string{"lemon", "mandarin", "cherries", "strawberry", "soursop", "guyava"}
	c := []string{"lemon", "mandarin", "cherries", "strawberry", "soursop", "guyava"}
	fmt.Println(equal(a, b), equal(a, c), equal(b, c))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
