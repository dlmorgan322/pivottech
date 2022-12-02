package main

import (
	"fmt"

	"github.com/dlmorgan322/pivot/calculator/calculator"
)

var x, y int
var o string

func main() {
	fmt.Println("Please select an operation:\n (+, -, *, or /)")
	fmt.Scanln(&o)

	if o == "+" {
		fmt.Println("Additon\n1st #:")
		fmt.Scan(&x)
		fmt.Println("2nd #:")
		fmt.Scan(&y)
		result := calculator.Add(x, y)
		fmt.Println(result)

	} else if o == "-" {
		fmt.Println("Subtraction\n1st #:")
		fmt.Scan(&x)
		fmt.Println("2nd #:")
		fmt.Scan(&y)
		result := calculator.Sub(x, y)
		fmt.Println(result)

	} else if o == "*" {
		fmt.Println("Multiplication\n1st #:")
		fmt.Scan(&x)
		fmt.Println("2nd #:")
		fmt.Scan(&y)
		result := calculator.Mult(x, y)
		fmt.Println(result)

	} else if o == "/" {
		fmt.Println("Division\n1st #:")
		fmt.Scan(&x)
		fmt.Println("2nd #:")
		fmt.Scan(&y)
		if x == 0 || y == 0 {
			fmt.Print("calculator: cannot divide by 0, please try again\n")
		} else {
			result := calculator.Div(x, y)
			fmt.Println(result)
		}
	} else {
		fmt.Println("Invalid Operation")
	}
}
