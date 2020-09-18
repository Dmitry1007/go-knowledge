package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	circle := circle{radius: 5}
	square := square{length: 10}
	circleArea := circle.area()
	squareArea := square.area()
	fmt.Println(circleArea)
	fmt.Println(squareArea)

	shapeCircleArea := info(circle)
	fmt.Println(shapeCircleArea)

	shapeSquareArea := info(square)
	fmt.Println(shapeSquareArea)
}
