package main

import "fmt"

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	fmt.Println("Writer Interface")
}