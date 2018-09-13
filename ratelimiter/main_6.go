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
	wg.Add(len(*b))

	// create rateLimiter buffered channel
	rateLimiter := make(chan bool, maxRate)

	for _, elem := range *b {

		// add any message to rate limiter
		rateLimiter <- true
	
		// execute with go routines
		go processElem(elem, &wg, rateLimiter)
	}

	wg.Wait()
}

// processElem process one batch element
func processElem(elem string, wg *sync.WaitGroup, rl <-chan bool) {
	defer wg.Done()

	fmt.Println("Processing element", elem)

	time.Sleep(time.Duration(500 + rand.Intn(500)) * time.Millisecond)

	<- rl
}

func main() {

	// process 5 docs batch
	for i := 1; i < 11; i++ {

		fmt.Println("\nLimiting max rate to ", i, "\n")
		create(10).process(i)
	}
}