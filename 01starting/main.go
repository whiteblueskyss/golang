package main

import "fmt"

func getValues(x, y int) (int, int) {
	return x+y, x*y
}

func main() {
	a := 10
	fmt.Println("Hello, World!", a)

	ok := false
	ok = true
	if ok==true {
		fmt.Println("You are pretty!")
	}

	var val int
	fmt.Println("Value is:", val)
	var x,y int
	// fmt.Scanln(&x) // to take input from user
	// fmt.Scanln(&y)
	x, y = 5, 7
	
	p, q := getValues(x,y)
	fmt.Println("Sum and Product:", p, q)
}