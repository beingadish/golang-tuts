package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views++
}

// ================ POST WITH MUTEX (POSTM) ==========================

type postM struct {
	views int
	mu    sync.Mutex
}

// Due to this Mutex Lock()/Unlock() Goroutines will wait till the Lock() gets released -> Unlocked()
func (p *postM) inc(wg *sync.WaitGroup) {
	defer func() {
		// Unlocking after Performing
		p.mu.Unlock()
		wg.Done()
	}()

	// Locking before Performing
	p.mu.Lock()
	p.views++
}

func main() {
	// GoLang tuts on Mutex
	var wg sync.WaitGroup
	myPost := post{views: 0}
	myPostM := postM{views: 0}

	for range 100 {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	for range 100 {
		wg.Add(1)
		go myPostM.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Without Mutex View Count (Results can vary) = ", myPost.views)
	fmt.Println("With Mutex View Count (Results will always = 100) = ", myPostM.views)
}
