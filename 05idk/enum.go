package main

import "fmt"

// enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

type PaymentStatus int

const (
	Pending PaymentStatus = iota // auto value will be 0 and next will be 1, 2, 3...
	Completed
	Failed
	Refunded
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func enumDemo() {
	changeOrderStatus(Prepared)
}
