package main

import "fmt"

func main() {
	// interfaceDemo()

	// deferDemo()

	// channelDemo()

	// goroutineDemo()

	// generic()

	// mutexDemo()

	// filesDemo()

	p := &Person{Name: "Selim"}

	SayHello(p)
}

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

// TODO: Add the method here so that *Person satisfies Greeter
func (p *Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}
func SayHello(g Greeter) {
	g.Greet()
}
