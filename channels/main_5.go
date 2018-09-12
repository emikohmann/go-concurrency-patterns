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

	ch := make(chan string, 4)

	// control go routine
	go func() {

		// read only one message from channel
		fmt.Println("Go routine say ", <- ch)
	} ()

	// go routines launcher
	for true {

		go ProcessUser(i, ch)

		time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)

		i++
	}
}

func ProcessUser(userID int, ch chan string) {

	// send data to non blocking channel
	ch <- fmt.Sprintf("I'm processing user %d", userID)

	fmt.Println("Message", userID, "was sent")
}