package main

import (
	"fmt"
)

func totalV(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func total(i []int) int {
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

// person is what it says it is bro
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	s := fmt.Sprintf(
		"My name is %v %v and I am %v years of age",
		p.first,
		p.last,
		p.age,
	)
	return s
}

func main() {
	xi := []int{1, 2, 3, 4}
	tv := totalV(xi...)
	fmt.Println(tv)

	i := []int{3, 4, 5}
	t := total(i)
	fmt.Println(t)

	d()

	p := person{
		first: "Shaq",
		last:  "Oneil",
		age:   42,
	}
	fmt.Println(p)
	sp := p.speak()
	fmt.Println(sp)
}
