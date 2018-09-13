package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Call a generator
	fravega := loadDataForUser("fravega")

	// Call new generator
	garbarino := loadDataForUser("garbarino")

	go func() {

		for true {

			fmt.Println("Generator say:", <- fravega)
			fmt.Println("Generator say:", <- garbarino)
		}	
	} ()

	time.Sleep(3 * time.Second)
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