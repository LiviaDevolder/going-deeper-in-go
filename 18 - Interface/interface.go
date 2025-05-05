package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height float64
	width  float64
}

func writeArea(f form) {
	fmt.Printf("The area is %0.2f", f.area())
}

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
