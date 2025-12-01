package main

import "fmt"

func sum(nums ...int) int { // nums is a variadic parameter of type int. It can take zero or more int arguments.
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func variadicDemo() {
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
