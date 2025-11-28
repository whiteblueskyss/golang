package main

import "fmt"

func loops() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			fmt.Println("before Jumping to lco")
			goto lco                 // jump to lco. Not recommended to use goto in code.
			fmt.Println("After lco") // this will never be executed because of the goto statement.
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("I am lco")
}
