package main

import (
	"fmt"
	"time"
)

// send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}

// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("processing...")
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	// emailChan receive only channel, done send only channel. if not specified, it's a bidirectional channel.

	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func channelDemo() {

	// messageChan := make(chan string)

	// messageChan <- "ping" // blocking. It will block here until some goroutine receives the data from the channel. But there is no other goroutine to receive the data. so it will cause deadlock.

	// msg := <-messageChan

	// fmt.Println(msg)
	// upper codes will cause deadlock. to fix it we can use goroutine. so it will use different thread and will be able to receive the data from the channel.

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100) // buffered channel with capacity of 100.
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending.")

	// this is important
	// close(emailChan)
	// <-done
	// emailChan <- "1@example.com" // if emailChan is a  buffered channel and it's not full, it will not cause panic / block. but if it's full, it will cause panic/block.

	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// here emailChan is receiving and sending data at the same time. so it will not block. if at any moment when emailChan is full, sending will block until some data is received from the channel.

	// done := make(chan bool)
	// go task(done)

	// <-done // block

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result // blocking

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "ping" // blocking

	// msg := <-messageChan

	// fmt.Println(msg)

}
