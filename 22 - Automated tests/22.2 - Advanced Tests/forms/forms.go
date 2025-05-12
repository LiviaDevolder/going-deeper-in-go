package forms

import (
	"fmt"
	"math"
)

type Rectangle struct {
	height float64
	width  float64
}

func WriteArea(f Form) {
	fmt.Printf("The area is %0.2f", f.Area())
}

type Circle struct {
	radius float64
}

type Form interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := Rectangle{10, 15}
	WriteArea(r)

	c := Circle{10}
	WriteArea(c)
}
