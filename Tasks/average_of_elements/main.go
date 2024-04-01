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

	var sum int

	for _, v := range arr {
		sum = sum + v
	}

	fmt.Println("average = ", (sum / len(arr)))
}
