package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20

	invertedNumber := invertSignal(number)

	fmt.Println(invertedNumber)
	fmt.Println(number)

	newNumber := 40

	fmt.Println(newNumber)
	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
