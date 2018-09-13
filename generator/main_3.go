package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Call a fan in function
	c := fanInMultiplex(
		loadDataForUser("fravega"),
		loadDataForUser("garbarino"),
	)

	go func() {

		for true {

			fmt.Println("Generators fan in say:", <- c)
		}	
	} ()

	time.Sleep(3 * time.Second)
}

// fanInMultiplex intercala la salida de ambos canales
func fanInMultiplex(input1, input2 chan string) chan string {
	c := make(chan string)

	go func() { for {c <- <- input1} } ()
	go func() { for {c <- <- input2} } ()

	return c
}

// Generator returns the channel
func loadDataForUser(username string) chan string {

	c := make(chan string)

	go func() {
		i := 0
		for true {
			c <- fmt.Sprintf("%s data %d", username, i)

    		time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)			
			
			i++
		}
	} ()

	return c
}