package main

import "fmt"

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2

	return sum, subtraction
}

func main() {
	result := sum(30, 1)

	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result2 := f("Function text")
	fmt.Println(result2)

	resultSum, _ := mathCalcs(10, 15)
	fmt.Println(resultSum)
}
