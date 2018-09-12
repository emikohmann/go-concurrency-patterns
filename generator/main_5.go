package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Call a fan in function
	c := FanInMultiplex(
		LoadDataForUser("fravega"),
		LoadDataForUser("garbarino"),
		LoadDataForUser("libertad"),
		LoadDataForUser("nike"),
		LoadDataForUser("adidas"),
		LoadDataForUser("sony"),
		LoadDataForUser("lg"),
		LoadDataForUser("mercadopago"),
	)

	go func() {

		for true {

			fmt.Println("fan in say", <- c)
		}	
	} ()

	time.Sleep(3 * time.Second)
}

func FanInMultiplex(inputs ...chan string) chan string {
	c := make(chan string)

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
func LoadDataForUser(username string) chan string {

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