package main

import "fmt"

var (
	d = 5
	b = 15
	c = 23
)

func main() {
	fmt.Println("\n1 + 1032 = ", 1+1032)
	fmt.Println("1.2 + 1032.7 = ", 1.2+1032.7)
	var x string = "Yo decidio eligir a tuyo"
	fmt.Println(x)
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	x = "first"
	var y string = " second"
	var z string = " third"
	fmt.Println(x + y + z)
	a := "Lolarey"
	fmt.Println(a)
	x = "May"
	fmt.Println("My dog's name is", x)
	dogName := "Max"
	fmt.Println("My dog's name is", dogName)
	f()
	fmt.Println("\n", b, c, d)
}

var dogsName = "Rex"

func f() {
	fmt.Println("That dog's name is", dogsName)
}
