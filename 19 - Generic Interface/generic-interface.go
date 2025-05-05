package main

import "fmt"

func generic(i interface{}) {
	fmt.Println(i)
}

func main() {
	generic("String")
	generic(1)
	generic(true)
}
