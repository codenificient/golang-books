package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Book!")
	echo()
	echo2()
	echo3()
	fmt.Println("End of Program")
}

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("echo", s)
}

func echo2 () {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		s = " "
	}
	fmt.Println("echo2", s)
}

func echo3() {
	fmt.Println("echo3", strings.Join(os.Args[1:], " "))
}