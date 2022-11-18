package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Go Book!")
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		s = " "
	}
	fmt.Println("echo2", s)
	fmt.Println("End of Program")
}
