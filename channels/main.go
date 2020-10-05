package main

import "fmt"

func main() {
	// Unbuffered bidirectional channel
	c1 := make(chan int)
	go func() {
		c1 <- 42
		c1 <- 43
	}()
	fmt.Println(<-c1, <-c1)

	// Buffered bidirectional channel
	// 1 is the buffer, meaning you can only place one value in the channel
	c2 := make(chan int, 1)
	c2 <- 57
	fmt.Println(<-c2)
	fmt.Printf("c2 is of Type %T \n", c2)

	// Directional channels
	// send only channel
	c3 := make(chan<- int)
	fmt.Printf("c3 is of Type %T \n", c3)

	// receive only channel
	c4 := make(<-chan int)
	fmt.Printf("c4 is of Type %T \n", c4)
}
