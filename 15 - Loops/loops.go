package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Increment i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Increment j", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Maria", "Fátima", "Lúcia"}

	for _, name := range names {
		fmt.Println(name)
	}

	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name":     "Leo",
		"lastname": "Valdez",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}
