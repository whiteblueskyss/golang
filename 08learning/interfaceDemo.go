package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailSender struct {
	Address string
}

type SMSSender struct {
	Number string
}

func (e EmailSender) Notify(message string) {
	fmt.Println("Sending EMAIL to", e.Address, ":", message)
}

func (s SMSSender) Notify(message string) {
	fmt.Println("Sending SMS to", s.Number, ":", message)
}

func SendNotification(n Notifier, msg string) {
	n.Notify(msg)
}

func interfaceDemo() {
	fmt.Println("=== Interface Demo ===")

	email := EmailSender{Address: "example@example.com"}
	sms := SMSSender{Number: "+1234567890"}

	SendNotification(email, "Hello via Email!")
	SendNotification(sms, "Hello via SMS!")
}
