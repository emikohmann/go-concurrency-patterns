package main

import (
    "fmt"
    "time"
    "math/rand"    
)

func main() {

	// simple go routine
	go processUsers()

	time.Sleep(3 * time.Second)
}

func processUsers() {

	i := 0

	for true {

		// ProcessUser execution is non blocking
		go processUser(i)

		time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)

		i++
	}
}

func processUser(userID int) {

	fmt.Println("I'm processing user", userID)
	
	// adding extra time in ProcessUser
	time.Sleep(2 * time.Second)
}