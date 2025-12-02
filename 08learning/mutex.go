package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // mutex to protect views field from concurrent access.
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1

	// p.mu.Lock()
	// p.views += 1
	// p.mu.Unlock()
}

func mutexDemo() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
