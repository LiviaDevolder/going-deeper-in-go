package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name     string `json:"name"`
	Pedigree string `json:"pedigree"`
	Age      uint   `json:"age"`
}

func main() {
	d := dog{"Rex", "Pug", 3}
	fmt.Println(d)

	jsonDog, err := json.Marshal(d)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonDog)

	fmt.Println(bytes.NewBuffer(jsonDog))

	d2 := map[string]string {
		"name": "Toby",
		"race": "Poodle",
	}

	jsonDog2, err2 := json.Marshal(d2)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(jsonDog2)

	fmt.Println(bytes.NewBuffer(jsonDog2))
}
