package main

import "fmt"

func structDemo() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	selim := User{"Selim", "selim@go.dev", true, 16}
	fmt.Println(selim)
	fmt.Printf("selim details are: %+v\n", selim)
	fmt.Printf("Name is %v and email is %v.", selim.Name, selim.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
