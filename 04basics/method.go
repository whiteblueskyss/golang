package main

import "fmt"

// Define a struct (type)
type Rectangle struct {
	width  float64
	height float64
}

// Method with VALUE receiver (a copy of Rectangle)
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method with POINTER receiver (can modify the original Rectangle)
func (r *Rectangle) Scale(factor float64) {
	r.width = r.width * factor
	r.height = r.height * factor
}

// Standalone FUNCTION (not a method)
func AreaFunction(r Rectangle) float64 {
	return r.width * r.height
}

func method() {
	rect := Rectangle{width: 10, height: 5}

	// Calling a METHOD with value receiver
	fmt.Println("Area (method):", rect.Area())

	// Calling the standalone FUNCTION
	fmt.Println("Area (function):", AreaFunction(rect))

	// Calling a METHOD with pointer receiver
	rect.Scale(2)
	fmt.Println("After scaling:", rect)

	// Area again after scaling
	fmt.Println("Area after scaling:", rect.Area())
}
