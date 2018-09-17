package main

import (
	"fmt"
	"time"
	"math/rand"
) 

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Duration(750 + rand.Intn(500)) * time.Millisecond)
		ch1 <- "I'm CH1!"
	} ()

	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Duration(750 + rand.Intn(500)) * time.Millisecond)
		ch2 <- "I'm CH2!"
	} ()

	// Wait for a communication
	select {

	case <-time.After(1000 * time.Millisecond):

		fmt.Println("Timeout")

	case msg1 := <- ch1:

		fmt.Println(msg1)

	case msg2 := <- ch2:

		fmt.Println(msg2)
	}

	fmt.Println("Program finished")
}