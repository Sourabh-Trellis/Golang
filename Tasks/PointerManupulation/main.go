package main

import "fmt"

// Task 1: Pointer Manipulation - Write a program that takes an integer input from the user and prints its memory address along with the value stored at that address.

func main() {

	var num int
	fmt.Println("Enter any number:")
	fmt.Scan(&num)

	fmt.Printf("Value at num : %v\n", num)
	fmt.Printf("Address of num : %v\n", &num)

}
