package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("Byte Counter")

	var c ByteCounter
	c.Write([]byte("hello there!"))
	fmt.Println(c)
}