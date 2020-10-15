package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (m mathError) Error() string {
	return fmt.Sprintf("a math error occured: %v %v %v", m.lat, m.long, m.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("Math: square root of negative number")
		me := fmt.Errorf("Math: square root of negative number: %v", f)
		return 0, mathError{"25.5678 N", "74.3490 W", me}
	}
	return 42, nil
}
