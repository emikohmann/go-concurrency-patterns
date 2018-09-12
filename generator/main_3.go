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
	)

	go func() {

		for true {

			fmt.Println("fan in say", <- c)
		}	
	} ()

	time.Sleep(3 * time.Second)
}

func FanInMultiplex(input1, input2 chan string) chan string {
	c := make(chan string)

	go func() { for {c <- <- input1} } ()
	go func() { for {c <- <- input2} } ()

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