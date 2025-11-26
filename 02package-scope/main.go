package main	

import (
	"fmt"
	"example.com/mathlib" // to import this package we need to create a module first using "go mod init example.com" command
)

func main() {
	x := mathlib.Add(5, 10)
	// x := add(5, 10) 
	// have to run both main.go and add.go at the same time command: go run main.go add.go

	fmt.Println("Sum is:", x)
}
