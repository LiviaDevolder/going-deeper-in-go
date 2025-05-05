package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	total := sum(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(total)

	write("Hello, world!", 1, 3, 4, 31)
}
