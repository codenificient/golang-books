package main

import (
	"fmt"
)

func main() {
	fmt.Println("Nonempty strings Golang")
	data := []string{"one", "", "three"}
	data1 := nonempty(data)
	fmt.Println(data1)
	fmt.Println("current state of data", data)
	data2 := nonempty2(data)
	fmt.Println(data2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
