package main

import (
	"fmt"
)

func sumV(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func sum(i []int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}

func d() {
	defer func() {
		fmt.Println("I'm so deffered in this func")
	}()
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}

func main() {
	xi := []int{1, 2, 3, 4}
	sv := sumV(xi...)
	fmt.Println(sv)

	i := []int{3, 4, 5}
	s := sum(i)
	fmt.Println(s)

	d()
}
