package main

import (
	"errors"
	"fmt"
)

func main() {
	// + INTEGERS
	var number int64 = -10000000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE
	var number3 rune = 10000000
	fmt.Println(number3)

	// BYTE = UINT8
	var number4 byte = 255
	fmt.Println(number4)
	// + END INTEGERS

	// + REAL NUMBERS

	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 123.45
	fmt.Println(realNumber2)

	realNumber3 := 123.56
	fmt.Println(realNumber3)

	// + END REAL NUMBERS

	// + STRINGS

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text 2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	// + END STRINGS

	// + BOOLEAN

	var boolean bool
	fmt.Println(boolean)

	// + END BOOLEAN

	// + ERROR

	var err error = errors.New("Intern error")
	fmt.Println(err)

	// + END ERROR
}
