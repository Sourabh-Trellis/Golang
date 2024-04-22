package main

import "fmt"

func main() {

	var dayOfWeek int
	fmt.Println("Enter day of week (1 for Monday, 2 for Tuesday.etc)")
	fmt.Scan(&dayOfWeek)

	switch dayOfWeek {
	case 1:
		fmt.Println("Day is Monday")
	case 2:
		fmt.Println("Day is Tuesday")
	case 3:
		fmt.Println("Day is Wednesday")
	case 4:
		fmt.Println("Day is Thursday")
	case 5:
		fmt.Println("Day is Friday")
	case 6:
		fmt.Println("Day is Saturday")
	case 7:
		fmt.Println("day is Sunday")
	}
}
