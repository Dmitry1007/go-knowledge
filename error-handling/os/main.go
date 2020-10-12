package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Create file
	f, err := os.Create("colors.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Red")
	// Write to file
	io.Copy(f, r)

	// Open file
	f, err = os.Open("colors.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// Read from file
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print read contents
	fmt.Println(string(bs))
}
