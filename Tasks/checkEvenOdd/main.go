package main

import "fmt"

// Task 5: Conditional Statements - Write a program that takes an integer input and checks if it is even or odd, then prints the result.

func main() {

	var num int
	fmt.Println("Enter any number :")
	fmt.Scan(&num)
	checkEvenOdd(num)

}

func checkEvenOdd(n int) {
	if n%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
