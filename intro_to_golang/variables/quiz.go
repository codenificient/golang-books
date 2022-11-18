package main

import "fmt"

func main() {
	fmt.Println("Converting Fahrenheit to Celsius")
	sevnty := F2C(70)
	fmt.Println("70 degrees F is ", sevnty, "degrees Celsius!")
}

func F2C(tmp float64) float64 {
	cel := (tmp - 32) * 5 / 9
	return cel
}
