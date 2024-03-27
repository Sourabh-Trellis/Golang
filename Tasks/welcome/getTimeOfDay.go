package main

import (
	"fmt"
	"time"
)

func getTimeOfDay() string {

	hrs, _ := fmt.Println(time.Now().Hour())
	fmt.Println(hrs)
	if hrs >= 5 && hrs < 12 {
		return "Morning"
	} else if hrs >= 12 && hrs < 17 {
		return "Afternoon"
	} else if hrs >= 17 && hrs < 21 {
		return "Evening"
	} else {
		return "Night"
	}

}
