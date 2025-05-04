package main

import "fmt"

func main() {
	// + ARITHMETICS

	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 10 * 5
	mod := 9 % 2

	fmt.Println(sum, subtraction, division, multiplication, mod)

	// + END ARITHMETICS

	// + ASSIGNMENT

	var var1 string = "String"
	var2 := "String 2"

	fmt.Println(var1, var2)

	// + END ASSIGNMENT

	// + RELATIONAL

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// + END RELATIONAL

	// + LOGICAL

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// + END LOGICAL

	// + UNARY

	number := 10

	number++
	number += 15

	fmt.Println(number)

	number--
	number -= 20

	number *= 3
	number /= 10
	number %= 3

	fmt.Println(number)

	// + END UNARY
}
