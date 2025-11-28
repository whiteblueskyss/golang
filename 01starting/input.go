package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //reads entire lin including space.
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T\n", input) // string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
