package functions

import (
	"fmt"
	"os"
	"time"
)

func CreateFile(name string) {
	fmt.Println("In mail")
	file, err := os.Create(name+".txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	go sendNotification(name)
	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")
	str := "Hello Sangameshwar,I am pleased to congratulate you on completing Milestone 1 of the Go language Training program.During the training sessions, we covered the following topics so far."
	file.WriteString(str)
	defer file.Close()
	file.WriteString(timeString)
}
