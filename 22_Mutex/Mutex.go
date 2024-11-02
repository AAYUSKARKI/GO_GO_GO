package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // Add a mutex field to the post struct
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock() // Lock the mutex before accessing the shared resource
	p.views++
	p.mu.Unlock() // Unlock the mutex after updating
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			myPost.incrementViews(&wg) // Only call incrementViews
		}()
	}
	wg.Wait()                 // Wait for all goroutines to finish
	fmt.Println(myPost.views) // Should correctly print 10
}
