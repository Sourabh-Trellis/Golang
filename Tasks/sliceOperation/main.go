package main

import "fmt"

// Task 2: Slice Operations - Write a program that takes a slice of integers as input and prints the sum of all the elements in the slice.

func main() {

	var size int
	fmt.Println("Enter size of slice :")
	fmt.Scan(&size)
	slice := make([]int, size)
	sum:=0
	
	for i := 1; i <= 5; i++ {
		var n int
		fmt.Println("Enter element ",i)
		fmt.Scan(&n)
		slice[i-1]=n
		sum = sum + n
	}

	fmt.Println("sum of all elements is ",sum)

}
