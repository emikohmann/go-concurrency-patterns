package main

import "fmt"

func main() {
	n := 100000

	leftmost := make(chan int)

	left, right := leftmost, leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)

		go gopher(left, right)

		left = right
	}

	go func(c chan int) {

		c <- 1
	} (right)

	fmt.Println(<- leftmost)
}

func gopher(left, right chan int) {
	left <- 1 + <- right
}