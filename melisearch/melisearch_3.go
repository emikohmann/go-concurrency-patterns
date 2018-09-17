package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Search func(query string) Result

type Result string

var (
	ClassiSearch = search("classi")
	ItemsSearch = search("items")
	HistorySearch = search("history")
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()
	
	results := MeliSearch("iphone")
	
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func MeliSearch(query string) []Result {
	timeout := time.After(190 * time.Millisecond)

	c := make(chan Result)
	results := make([]Result, 0)

	go func() {c <- ClassiSearch(query)} ()
	go func() {c <- ItemsSearch(query)} ()
	go func() {c <- HistorySearch(query)} ()

	for i := 0; i < 3; i++ {
		select {

		case result := <- c:
				results = append(results, result)

		case <-timeout:
				fmt.Println("Timed out")
				return results
		}
	}

	return results
}

func search(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		return Result(fmt.Sprintf("\n%s result for %s", kind, query))
	}
}