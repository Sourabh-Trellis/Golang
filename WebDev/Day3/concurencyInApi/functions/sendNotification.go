package functions

import (
	"fmt"
	"os"
	"time"
)

func sendNotification(name string) {
	fmt.Println("in Notify")
	file, err := os.Create(name + "_notify.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")
	str := "Training Progress Review:Overall, I am pleased to report that Sangam has made significant progress in his Golang training. He has demonstrated a solid grasp of the foundational concepts and have shown enthusiasm for further learning and development."
	file.WriteString(str)
	file.WriteString(timeString)
	defer file.Close()
}
