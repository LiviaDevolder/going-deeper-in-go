package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      uint8
	height   uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"John", "Doe", 30, 165}
	fmt.Println(p1)

	s1 := student{p1, "IT", "UFF"}
	fmt.Println(s1)
}
