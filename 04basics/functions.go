
package main

import "fmt"

func add(a, b int) int { // a, b both are int here
	return a + b
}

func processIt(fn func(a int) int) { // function as parameter. Here fn is a function that takes int parameter and returns int
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		 return 4
	}
}

func functiomDemo() {
    // Call processIt with an inline function that prints its argument and returns it
    processIt(func(a int) int {
	 fmt.Println("processIt called with:", a)
	 return a
    })

    // Call processIt2, which returns a function, then call that function and print the result
    fn := processIt2()
    result := fn(10)
    fmt.Println("processIt2 returned function result:", result)
}
