package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of array")
	fmt.Scan(&size)
	arr := make([]int, size)

	fmt.Printf("Enter %v elements in the array", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	var max int

	for _, v := range arr {
		if v > max {
			max=v
		}
	}
	fmt.Println("max = ",max)
}
