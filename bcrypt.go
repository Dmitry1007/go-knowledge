package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(bs)

	pass1 := `password124`
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass1))
	if err != nil {
		fmt.Println(err)
	}
}
