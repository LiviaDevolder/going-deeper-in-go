package main

import (
	"fmt"
	"test-intro/address"
)

func main() {
	addressType := address.AddressType("Road 66")

	fmt.Println(addressType)
}
