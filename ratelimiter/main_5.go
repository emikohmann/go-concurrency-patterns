package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type batch []string

// load generate random batch
func create(n int) *batch {
	documents := make(batch, 0)
	for i := 0; i < n; i++ {
		documents = append(
			documents,
			fmt.Sprintf(
				"document-%d", 100 + rand.Intn(200),
			),
		)
	}
	return &documents
}

// process iterates batch elems
func (b *batch) process(maxRate int) {

	fmt.Println("Processing batch", &b)

	// use waitgroup for execute all go routines
	var wg sync.WaitGroup

	// set go routines count
	wg.Add(len(*b))

	// create rateLimiter buffered channel
	rateLimiter := make(chan bool, maxRate)

	for _, elem := range *b {

		// add any message to rate limiter
		rateLimiter <- true
	
		// execute with go routines
		go processElem(elem, &wg, rateLimiter)
	}

	// wait for all go routines
	wg.Wait()
}

// processElem process one batch element
func processElem(elem string, wg *sync.WaitGroup, rl <-chan bool) {
	// decrement waitgroup counter
	defer wg.Done()

	fmt.Println("Processing element", elem)

	time.Sleep(time.Duration(500 + rand.Intn(500)) * time.Millisecond)

	// free a place

	fmt.Println("\tElement finished", elem)
	<- rl
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	// process 10 docs batch
	// adding rate limit by arg
	create(10).process(3)
}