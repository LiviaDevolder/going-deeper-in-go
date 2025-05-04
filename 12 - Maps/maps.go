package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Livia",
		"Devolder": "Devolder",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Livia",
			"last":  "Devolder",
		},
		"course": {
			"name":   "IT",
			"campus": "campus 1",
		},
	}

	fmt.Println(user2)
	delete(user2, "course")
	fmt.Println(user2)
}
