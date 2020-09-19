package main

import (
	"crypto/sha256"
	"fmt"
)

func makeTree(s []string) []string {
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

	return level3
}

func hashPuppy(s []string) []string {
	var hashedValues []string
	for _, v := range s {
		h := sha256.New()
		h.Write([]byte(v))
		hashed := fmt.Sprintf("%x", h.Sum(nil))
		hashedValues = append(hashedValues, hashed)
	}
	return hashedValues
}

func main() {
	txs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	makeTree(txs)

	hashed := hashPuppy(txs)
	fmt.Println(hashed)

	h := sha256.New()
	h.Write([]byte(txs[0]))
	hashy := fmt.Sprintf("\n%x", h.Sum(nil))
	fmt.Println(hashy)
}
