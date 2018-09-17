package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	// creamos canales de ejemplo
	messages := make(chan string)
	quit := make(chan bool)

	// lanzamos una go routine que carga mensajes
	go func() {
		for i := 0; ; i++ {
			messages <- fmt.Sprintf("Hello %d", i)
			time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)
			i++
		}
	}()

	// lanzamos una go routine que dispara el finish
	go func() {
		time.Sleep(5000 * time.Millisecond)
		quit <- true
	}()

	// ejecutamos segun el caso disponible
	FOR:
	for {
		select {
	
			case msg := <- messages: 
		
					fmt.Println("Received message from c1 channel", msg)
		
			case <- quit:
		
					fmt.Println("Select finished")
					break FOR
			}
	}

	fmt.Println("Program finished")
}