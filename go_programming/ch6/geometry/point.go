package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance1(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
	
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance1(path[i])
		}
	}
	return sum
}

func main() {
	fmt.Println("Geometry")
}