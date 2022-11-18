package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	dob       time.Time
	Position  string
	Salary    float32
	ManagerID int
}

func main() {
	fmt.Println("Structures in Go")

	var dilbert Employee
	dilbert.Name = "Dilbert"
	dilbert.Address = "123 Coding Lane"
	dilbert.Position = "Software Developer"
	dilbert.dob = time.Now()
	dilbert.Salary = 47000.3
	fmt.Println(dilbert)
}
