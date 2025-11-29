package main

import (
	"fmt"
	"io"
	"os"
)

func fileDemo() {
	fmt.Println("Welcome to files in golang")
	content := "This words needs to go in a txt file which will automatically created by golang code."

	file, err := os.Create("./mytxtfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mytxtfile.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
