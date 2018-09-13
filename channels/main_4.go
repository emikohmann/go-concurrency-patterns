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

	// buffered non blocking channel
	ch := make(chan string, 4)

	// control go routine
	go func() {

		// no read message
	} ()

	// go routines launcher
	for true {

		go processUser(i, ch)

		time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)

		i++
	}
}

func processUser(userID int, ch chan string) {

	// send data to non blocking channel
	ch <- fmt.Sprintf("I'm processing user %d", userID)

	fmt.Println("Message", userID, "was sent")
}