package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Call a generator
	fravega := LoadDataForUser("fravega")

	go func() {

		for true {

			fmt.Println("fravega say", <- fravega)
		}	
	} ()

	time.Sleep(3 * time.Second)
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