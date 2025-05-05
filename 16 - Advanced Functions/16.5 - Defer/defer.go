package main

import "fmt"

func func1() {
	fmt.Println("Function 1")
}

func func2() {
	fmt.Println("Function 2")
}

func main() {
	defer func1()

	func2()
}
