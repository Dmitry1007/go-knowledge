package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func main() {
	i := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].first < people[j].first
	})
	fmt.Println("By First Name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("By Age:", people)
}
