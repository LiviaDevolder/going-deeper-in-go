package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less or equal 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Number greater then 0")
	}
}
