package main

import (
    "fmt"
    "time"
    "math/rand"    
)

func main() {

	// simple go routine
	go ProcessUsers()

	time.Sleep(3 * time.Second)
}

func ProcessUsers() {

	i := 0

	for true {

		// ProcessUser execution is non blocking
		go ProcessUser(i)

		time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)

		i++
	}
}

func ProcessUser(userID int) {

	fmt.Println("I'm processing user", userID)
	
	// adding extra time in ProcessUser
	time.Sleep(2 * time.Second)
}