package main

import (
	"fmt"
)

// We are pointing to an address in memory where a type int is located
// We are passing the value of an address
func point(y *int) {
	fmt.Println("y =", y)
	fmt.Println("*y =", *y)
	// Now we are derefering by taking the address and pointing to the value
	*y = 55
	fmt.Println("y =", y)
	fmt.Println("*y =", *y)
}

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

	x := 0
	fmt.Println("x =", x)
	fmt.Println("&x =", &x)
	point(&x)
	fmt.Println("x =", x)
	fmt.Println("&x =", &x)
}
