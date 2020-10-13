package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error Happened", err)

		// Fatal is equivalent to Println() followed by os.Exit(1)
		// Entire program is going to shut down immediately!
		// log.Fatalln(err)

		// Panic is equivalent to Println() followed by a cal to panic()
		// log.Panicln(err)

		// Panic function call panic after writing the log message
		// panic(err)
	}
	defer f2.Close()

	fmt.Println("Check the log.txt file")
}
