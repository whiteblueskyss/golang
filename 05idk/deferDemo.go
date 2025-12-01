package main

import "fmt"

func deferDemo() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

	// defer statements are executed in LIFO (Last In First Out) order. Like a stack.
}

// world, One, Two
// 0, 1, 2, 3, 4
// hello, 43210, two, One, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// defer will execute at the end of the function in LIFO order.

// Example uses:

// Close a file

// Close a database connection

// Unlock a mutex

// Recover from panic

// Logging end of function
