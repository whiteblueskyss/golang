package main

import (
	"bufio" // used for buffered input and output
	"fmt"
	"os"      // used for operating system functionality like reading from stdin, writing to stdout, and working with files.
	"strconv" // used for string conversions. methods like Atoi, ParseFloat, itoa and more.
	"strings" // used for string manipulation like trimming, splitting, replacing, and more.
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
