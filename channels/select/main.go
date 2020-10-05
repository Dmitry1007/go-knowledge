package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("Bout To Exit")
}

func send(even, odd, quit chan<- int) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case val := <-even:
			fmt.Println("from even channel \t", val)
		case val := <-odd:
			fmt.Println("from odd channel \t", val)
		case val := <-quit:
			fmt.Println("from quit channel \t", val)
			return
		}
	}
}
