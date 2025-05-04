package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Struct file")

	var u user
	u.name = "Livia"
	u.age = 22
	fmt.Println(u)

	addressExample := address{"street", 0}

	u2 := user{
		"Livia", 22, addressExample,
	}
	fmt.Println(u2)

	u3 := user{age: 22}
	fmt.Println(u3)
}
