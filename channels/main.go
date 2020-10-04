package main

import "fmt"

func main() {
	// unbuffered channel
	c1 := make(chan int)
	go func() {
		c1 <- 42
		c1 <- 43
	}()
	fmt.Println(<-c1, <-c1)

	// buffered channel
	// 1 is the buffer, meaning you can only place one value in the channel
	c2 := make(chan int, 1)
	c2 <- 57
	fmt.Println(<-c2)

}
