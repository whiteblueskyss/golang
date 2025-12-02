package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) { // T any,  T interface{} (same to any) , T string | int | bool we can also use like this.

	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// LIFO
type stack[T any] struct {
	elements []T
}

func generic() {
	myStack := stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(myStack)

	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	// printStringSlice(names)
	printSlice(vals, "john")
}
