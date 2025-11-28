package main

import (
	"fmt"
	"mymath/mathlib" // importing the custom mathlib package
)

func main() {
	x := mathlib.Add(5, 10)
	y := mathlib.Sub(10, 3)

	fmt.Println("Sum is:", x)
	fmt.Println("Difference is:", y)
}
