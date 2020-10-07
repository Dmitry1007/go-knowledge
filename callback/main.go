package main

import (
	"fmt"
)

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func main() {
	i := []int{1, 2, 3, 4, 5}
	s := sum(i...)
	fmt.Println(s)

	e := even(sum, i...)
	fmt.Println(e)

	o := odd(sum, i...)
	fmt.Println(o)
}
