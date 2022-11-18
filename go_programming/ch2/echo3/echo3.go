package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Book!")
	fmt.Println("echo3", strings.Join(os.Args[1:], " "))
	fmt.Println("End of Program")
}
