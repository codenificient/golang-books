package main

import "fmt"

type Var string

type literal float64

type unary struct {
	op rune
	x Expr
}

type binary struct {
	op rune
	x, y Expr
}

type call struct {
	fn string
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
}

func main() {
	fmt.Println("Evaluations")
}