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
		loadDataForUser("libertad"),
		loadDataForUser("nike"),
		loadDataForUser("adidas"),
		loadDataForUser("sony"),
		loadDataForUser("lg"),
		loadDataForUser("mercadopago"),
	)

	go func() {

		for true {

			fmt.Println("Generators fan in say:", <- c)
		}	
	} ()

	time.Sleep(3 * time.Second)
}

// fanInMultiplex intercala la salida de n canales
func fanInMultiplex(inputs ...chan string) chan string {
	c := make(chan string)

	// inputs es un parámetro variádico
	for _, in := range inputs {
		go func(current chan string) {
			for {
				c <- <- current
			}
		} (in)
	}

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