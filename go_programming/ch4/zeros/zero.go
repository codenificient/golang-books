package main

import "fmt"

func main() {
	fmt.Println("Zeros in Golang Array")

}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}