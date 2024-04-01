package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----To-Do-List App-----")
	var toDoList = make(map[int]string)
	var count int = 0
	for {
		fmt.Println("1)Add item to to-do-lost ")
		fmt.Println("2)delete item from to-do-list")
		fmt.Println("3)list all to-do items")
		fmt.Println("4)Exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Print("Enter task you want to add :")
			fmt.Scan(&task)
			task = strings.TrimSpace(task)
			toDoList[count] = task
			count++
		case 2:
			fmt.Print("Enter number of the task you want to delete:")
			var num int
			fmt.Scan(&num)
			delete(toDoList, num)
			fmt.Println("Deleted Successfully")
		case 3:
			fmt.Println("Pending tasks are")
			for k, v := range toDoList {
				fmt.Printf("%v := %v\n", k, v)
			}
		case 4:
			fmt.Println("-----Thank you-----")
			return
		}
	}

}
