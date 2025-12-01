package main

import "fmt"

func getValues(x, y int) (int, int) {
	return x + y, x * y
}

func basicsDemo() {
	fmt.Println("=== Basics Demo ===")

	a := 10
	fmt.Println("Hello, World!", a)

	ok := false
	ok = true
	if ok == true {
		fmt.Println("You are pretty!")
	}

	var val int
	fmt.Println("Value is:", val)
	var x, y int
	// fmt.Scanln(&x) // to take input from user
	// fmt.Scanln(&y)
	x, y = 5, 7

	p, q := getValues(x, y)
	fmt.Println("Sum and Product:", p, q)
}

// we can declare a variable inside if construct
// if age := 20; age >= 18 {
// 	fmt.Println("person is an adult", age)
// } else if age >= 12 {
// 	fmt.Println("person is teenager", age)
// }

// go does not have ternary, you will have to use normal if else
