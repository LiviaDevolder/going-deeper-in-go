package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var arr1 [5]int
	arr1[0] = 10
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr3)

	slc := []int{2, 4, 6, 8}
	fmt.Println(slc)

	fmt.Println(reflect.TypeOf(slc))
	fmt.Println(reflect.TypeOf(arr1))

	slc = append(slc, 10)
	fmt.Println(slc)

	slc2 := arr2[1:3]
	fmt.Println(slc2)

	arr2[1] = 256
	fmt.Println(slc2)
}
