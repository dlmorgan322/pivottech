package calculator

import "math"

//test

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
