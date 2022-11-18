package main

import (
	"crypto/sha256"
	"fmt"
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}


func main() {
	fmt.Println("Golang Zeros")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	a := [32]byte{3, 2, 31}
	p1 := &a

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(a)
	fmt.Println(zero(*p1))
	// fmt.Println(zero(*c2))
}
