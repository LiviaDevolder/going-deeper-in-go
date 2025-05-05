package main

import "fmt"

var n int

func init() {
	fmt.Println("Running init function")
	n = 10
}

func main() {
	fmt.Println("Running main function", n)
}
