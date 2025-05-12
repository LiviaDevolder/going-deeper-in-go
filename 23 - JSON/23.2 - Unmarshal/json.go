package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name     string `json:"-"`
	Pedigree string `json:"pedigree"`
	Age      uint   `json:"age"`
}

func main() {
	jsonDog := `{"name":"Rex","pedigree":"Pug","age":3}`

	var d dog

	if err := json.Unmarshal([]byte(jsonDog), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	jsonDog2 := `{"name":"Toby","pedigree":"Poodle"}`

	d2 := make(map[string]string)

	if err := json.Unmarshal([]byte(jsonDog2), &d2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d2)
}
