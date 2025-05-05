package main

import "fmt"

func main() {
	result := func(text string) string {
		return fmt.Sprintf("got -> %s", text)
	}("Hello, world!")

	fmt.Println(result)
}
