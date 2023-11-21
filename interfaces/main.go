package main

import (
	"fmt"
	"math"
)

type square struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

func printArea(f form) {
	fmt.Printf("The area is %0.2f", f.area())
}

func (s square) area() float64 {
	return s.height * s.width
}

func (c circle) area() float64 {
	return math.Pi + (c.radius * c.radius)
}

func main() {

	s := square{20, 20}
	printArea(s)

	c := circle{20}
	printArea(c)
}
