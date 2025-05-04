package main

import "fmt"

func main() {
	slc := make([]float32, 10, 11)
	fmt.Println(slc)

	slc = append(slc, 5)
	slc = append(slc, 6)

	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slc2 := make([]float32, 5)
	fmt.Println(slc2)
	slc2 = append(slc2, 10)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))
}
