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

	var element_to_search int
	fmt.Println("Enter element you want to search :")
	fmt.Scan(&element_to_search)

	for idx, v := range arr {
		if v == element_to_search {
			fmt.Println("element found at index ", idx)
			return
		}
	}
	fmt.Println("Element not found.")
}
