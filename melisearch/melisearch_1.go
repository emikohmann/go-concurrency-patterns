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
	results := make([]Result, 0)
	results = append(results, ClassiSearch(query))
	results = append(results, ItemsSearch(query))
	results = append(results, HistorySearch(query))
	return results
}

func search(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		return Result(fmt.Sprintf("\n%s result for %s", kind, query))
	}
}