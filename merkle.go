package main

import (
	"fmt"
)

func merk(s []string) {
	fmt.Println("Level 0 ", s)
	var level1 []string
	for i, v := range s {
		if i%2 == 0 {
			level1 = append(level1, v+s[i+1])
		}
	}
	fmt.Println("Level 1 ", level1)

	var level2 []string
	for i, v := range level1 {
		if i%2 == 0 {
			level2 = append(level2, v+level1[i+1])
		}
	}
	fmt.Println("Level 2 ", level2)

	var level3 []string
	for i, v := range level2 {
		if i%2 == 0 {
			level3 = append(level3, v+level2[i+1])
		}
	}
	fmt.Println("Level 3 ", level3)
}

func main() {
	txs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	merk(txs)
}
