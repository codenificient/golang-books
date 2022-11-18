package main

import "fmt"

func main() {
	fmt.Println("\nNumbers..")
	fmt.Println("1 + 2 =", 1+2)

	fmt.Println("\nStrings..")
	fmt.Println(len("Welcome to Golang :)"))
	fmt.Println("Hello to my Gophers"[0])

	fmt.Println("\nBooleans..")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
