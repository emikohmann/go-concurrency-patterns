package main

import (
	"fmt"
	"time"
	"math/rand"
) 

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	g := GetData("fravega")

	// Wait for a communication
	select {

	case <-time.After(1000 * time.Millisecond):

		fmt.Println("Timeout")

	case msg := <- g:

		fmt.Println(msg)
	}

	fmt.Println("Program finished")
}

// Get data returns the channel
func GetData(user string) chan string {
	ch := make(chan string)

	go func() {

		time.Sleep(time.Duration(750 + rand.Intn(500)) * time.Millisecond)

		ch <- fmt.Sprintf("%s is done", user)
	} ()

	return ch
}