// Student Grade Tracker:
// Build an application to track student grades for multiple subjects.
// Allow users to add grades, calculate average grades, and determine the overall grade.
// package main

// import (
// 	"fmt"
// )

// type emptyMap map[string]int

// type student struct {
// 	name   string
// 	rollNo int
// 	grades emptyMap
// }

// func main() {
// 	fmt.Println("-------Grade Tracker-------")
// 	var studentList = make( map[int]student)
// 	var counter int = 1
// 	for {
// 		fmt.Println("1)Add student.")
// 		fmt.Println("2)Grades.")
// 		fmt.Println("3)View grades list.")
// 		fmt.Println("3)Exit")
// 		fmt.Println("Choose option")
// 		var option int
// 		fmt.Scan(&option)

// 		switch option {

// 		case 1:
// 			var name string
// 			fmt.Print("Enter Student name:")
// 			fmt.Scan(&name)

// 			student := student{}
// 			student.name=name
// 			student.rollNo=counter
// 			student.grades= make(emptyMap)

// 			studentList[student.rollNo]=student
// 			// fmt.Println(studentList)
// 		case 2:
// 			fmt.Print("Enter student roll number :")
// 			var rollNumber int
// 			fmt.Scan(&rollNumber)

// 		case 3:

// 		}
// 	}
// }

package main

import "fmt"

func main() {
	var name string
	var Physics int
	var chemistry int
	var mathematics int
	var socialStudies int

	fmt.Println("----------------------------")

	fmt.Print("Enter your name :")
	fmt.Scan(&name)

	for {

		fmt.Println("----------------------------")

		fmt.Println("1)Add grades.")
		fmt.Println("2)Calculate average.")
		fmt.Println("3)Determine overall grades")
		fmt.Println("4)Exit")
		fmt.Println("Choose the option :")
		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("\nEnter your marks")

			fmt.Print("Physics :")
			fmt.Scan(&Physics)

			fmt.Print("\nChemistry :")
			fmt.Scan(&chemistry)

			fmt.Print("\nMathematics :")
			fmt.Scan(&mathematics)

			fmt.Print("\nSocial Studies :")
			fmt.Scan(&socialStudies)
		case 2:
			var average float32
			average = (float32(Physics) + float32(chemistry) + float32(mathematics) + float32(socialStudies)) / 4
			fmt.Println("----------------------------")
			fmt.Println("Average is ", average)
		case 3:
			average := (float32(Physics) + float32(chemistry) + float32(mathematics) + float32(socialStudies)) / 4
			fmt.Println("----------------------------")
			if average > 90 {
				fmt.Println("You achieved \"A+\" grade.")
			} else if average <= 90 && average > 80 {
				fmt.Println("You achieved \"A\" grade")
			} else if average <= 80 && average > 70 {
				fmt.Println("You achieved \"B+\" grade")
			} else if average <= 70 && average > 60 {
				fmt.Println("You achieved \"B\" grade")
			} else if average <= 60 && average > 50 {
				fmt.Println("You achieved \"C+\" grade")
			} else if average <= 50 && average > 40 {
				fmt.Println("You achieved \"C\" grade")
			} else {
				fmt.Println("Fail!!!!!!!")
			}
		case 4:
			fmt.Println("----------Thank you----------")
			return
		}
	}
}
