package main

import (
	"fmt"
	"module/auxiliary"
)

func main() {
	fmt.Println("Writing from main package")
	auxiliary.Write()
}
