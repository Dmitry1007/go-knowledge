package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
}
