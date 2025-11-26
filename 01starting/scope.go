// global scope, local, package
// block scope, function scope, shadowing

package main

import "fmt"

// Global variable - accessible everywhere in the package
var globalVar = "I'm global"

func scopeDemo() {
	// Local variable in scopeDemo function
	localVar := "I'm local to scopeDemo"

	fmt.Println("=== Scope Demo ===")
	fmt.Println("Global variable:", globalVar)
	fmt.Println("Local variable:", localVar)

	// Block scope
	if true {
		blockVar := "I'm only in this block"
		fmt.Println("\nInside if block:", blockVar)
		fmt.Println("Can access local:", localVar)
	}
	// fmt.Println(blockVar) // This would cause an error - blockVar not accessible here

	// Function scope
	demonstrateFunction()

	// Shadowing example
	x := 10
	fmt.Println("\nOuter x:", x)

	if true {
		x := 20 // This shadows the outer x
		fmt.Println("Inner x (shadowed):", x)
	}

	fmt.Println("Outer x again:", x)
}

func demonstrateFunction() {
	functionVar := "I'm local to demonstrateFunction"
	fmt.Println("\nInside function:", functionVar)
	fmt.Println("Can access global:", globalVar)
	// fmt.Println(localVar) // This would cause an error - localVar not accessible here
}
