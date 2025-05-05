package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Succefuly recovered")
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("AVERAGE IS EXECTLY 6")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
}
