package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateFormat("liviadevolder@gmail.com")
	fmt.Println(err)
}
