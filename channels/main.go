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

	c5 := make(chan int)
	go sendChan(c5, 45)
	// receiveChan will block until sendChan does its thang
	receiveChan(c5)

	c6 := make(chan int)
	// send
	go func() {
		for i := 0; i <= 10; i++ {
			c6 <- i
		}
		close(c6)
	}()

	// receive
	// will block and continue to loop until the channel is closed
	for v := range c6 {
		fmt.Println(v)
	}

	fmt.Println("Bout To Exit")
}

func sendChan(c chan<- int, val int) {
	c <- val
}

func receiveChan(c <-chan int) {
	fmt.Println("receiveChan function", <-c)
}
