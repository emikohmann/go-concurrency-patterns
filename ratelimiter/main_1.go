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

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	// show 10 docs batch
	fmt.Println(*create(10))
}