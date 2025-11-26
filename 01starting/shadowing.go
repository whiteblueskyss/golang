package main

import "fmt"

func shadowingDemo() {
	fmt.Println("=== Shadowing Demo ===")

	// Shadows package-level msg with a new local variable
	msg := "main-level"
	fmt.Println("msg in main:", msg)

	// Block-level shadowing
	x := 10
	fmt.Println("outer x:", x)
	{
		x := 20 // shadows outer x inside this block
		fmt.Println("inner x:", x)
	}
	fmt.Println("after block x:", x)

	// Short variable declaration can shadow an existing name when at least one name is new
	y := 1
	fmt.Println("outer y:", y)
	y, z := 2, 3 // y is shadowed, z is a newly declared variable
	fmt.Println("shadowed y:", y, "z:", z)

	// Shadowing a function parameter by introducing a new variable with the same name in an inner block
	show := func(a int) {
		fmt.Println("param a:", a)
		{
			a := a + 1 // shadows parameter a inside this inner block
			fmt.Println("shadowed a:", a)
		}
		fmt.Println("after inner block a:", a)
	}
	show(5)
}
