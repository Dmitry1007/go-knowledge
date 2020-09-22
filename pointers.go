package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("&a gives you the address of a in memory ", &a)
	fmt.Printf("The Type of a is %T \n", a)
	fmt.Printf("The Type of &a is %T \n", &a)

	b := &a
	fmt.Println("b now points to the same address as a ", b)
	fmt.Println("*b gives you the value stored at the address ", *b)

	*b = 123
	fmt.Println("changing the value at the address of a means changing a's value ", a)
}
