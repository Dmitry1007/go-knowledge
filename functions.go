package main

import (
	"fmt"
)

func foo(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	xi := []int{1, 2, 3, 4}
	f := foo(xi...)
	fmt.Println(f)

	i := []int{3, 4, 5}
	b := bar(i)
	fmt.Println(b)
}
