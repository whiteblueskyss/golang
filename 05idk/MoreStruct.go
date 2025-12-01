package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition - embedding one struct within another
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// you can access order.name or order.customer.name directly if you want.

func newOrder(id string, amount float32, status string) *order { // like constructor. go doesn't have constructors. so we create functions that return struct pointers.

	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) { // pointer receiver to modify the original struct
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func moreStructs() {
	newCustomer := customer{
		name:  "john",
		phone: "1234567890",
	}
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		// customer: customer{
		// 	name:  "john",
		// 	phone: "1234567890",
		// },
		customer: newCustomer,
	}

	newOrder.customer.name = "robin"
	fmt.Println(newOrder)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true} // anonymous struct literal assignment to variable language. It can't be reused. Annoymous structs are useful for one-time use cases. like testing or temporary data structures.

	// fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)
}
