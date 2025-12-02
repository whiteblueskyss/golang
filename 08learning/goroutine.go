package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Mark this task as done when the function exits. Defer ensures this runs even if there's a panic. Defer alaways executes at the end of the function.

	fmt.Println("doing task", id)
}

func goroutineDemo() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish before exiting the main function. This blocks until the WaitGroup counter is zero. Best practice to write at the end of the main function.
}
