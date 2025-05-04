package main

import "fmt"

func weekday(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid value"
	}
}

func weekdayFallthrough(number int) string {
	var weekday string

	switch number {
	case 1:
		weekday = "Sunday"
		fallthrough
	case 2:
		weekday = "Monday"
	case 3:
		weekday = "Tuesday"
	case 4:
		weekday = "Wednesday"
	case 5:
		weekday = "Thursday"
	case 6:
		weekday = "Friday"
	case 7:
		weekday = "Saturday"
	default:
		weekday = "Invalid value"
	}

	return weekday
}

func main() {
	fmt.Println(weekday(4))
	fmt.Println(weekday(10))

	fmt.Println(weekdayFallthrough(1))
}
