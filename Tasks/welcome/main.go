package main

import (
	"fmt"
	"time"
)

func main() {

	var Name string
	fmt.Print("Enter your name:")

	fmt.Scan(&Name)
	var timeOfDay string
	hrs := time.Now().Hour()
	// fmt.Println(hrs)
	if hrs >= 5 && hrs < 12 {
		timeOfDay = "Morning"
	} else if hrs >= 12 && hrs < 17 {
		timeOfDay = "Afternoon"
	} else if hrs >= 17 && hrs < 21 {
		timeOfDay = "Evening"
	} else {
		timeOfDay = "Night"
	}
	fmt.Printf("Good %v, %v.\n",timeOfDay,Name)
}
