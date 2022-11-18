package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

func main() {
	fmt.Println("Reader Interface")
}
