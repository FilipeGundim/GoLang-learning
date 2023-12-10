package Forms

import (
	"fmt"
	"math"
)

type Square struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Form interface {
	Area() float64
}

func PrintArea(f Form) {
	fmt.Printf("The area is %0.2f", f.Area())
}

func (s Square) Area() float64 {
	return s.Height * s.Width
}

func (c Circle) Area() float64 {
	return math.Pi + (c.Radius * c.Radius)
}

func main() {

	s := Square{20, 20}
	PrintArea(s)

	c := Circle{20}
	PrintArea(c)
}
