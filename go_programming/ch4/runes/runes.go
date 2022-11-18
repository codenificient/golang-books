package main

import "fmt"

func main() {
	fmt.Println("Runes in Go")

	var runes []rune
	for _, r := range "Hello, World" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
