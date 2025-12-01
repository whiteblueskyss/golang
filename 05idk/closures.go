package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1 // this inner function increments and returns the count variable from the outer function's scope. it will remember the value of count between calls.
		return count
	}

}

func closuresDemo() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	// Each call to increment() will remember the updated count value from the previous call.
}
