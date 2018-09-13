package main

import (
	"fmt"
	"math/rand"
	"time"
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
func (b *batch) process() {

	fmt.Println("Processing batch", &b)

	for _, elem := range *b {
	
		processElem(elem)
	}
}

// processElem process one batch element
func processElem(elem string) {

	fmt.Println("Processing element", elem)

	time.Sleep(time.Duration(500 + rand.Intn(500)) * time.Millisecond)
}

func main() {

	// process 5 docs batch
	create(10).process()
}