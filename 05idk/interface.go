package main

import "fmt"

type paymenter interface {
	pay(amount float32) // if need to return something, we can define return type here. like pay(amount float32) bool
	refund(amount float32, account string)
}

// If an interface has multiple methods, your struct must implement all of them to satisfy the interface.
// If you implement only one method (and the interface has more), your struct will not satisfy the interface. so it will cause a compile-time error.

type payment struct {
	gateway paymenter
}

// Open close principle - software entities (classes, modules, functions, etc.) should be open for extension but closed for modification.

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) { // no need to tell that it implements paymenter interface. go does it automatically if the methods match.

	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

func interfaceDemo2() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}

// dependency inversion principle - high level modules should not depend on low level modules. both should depend on abstractions (interfaces).
