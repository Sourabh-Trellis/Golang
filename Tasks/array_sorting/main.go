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

	for i:=0 ; i<len(arr)-1 ; i++ {
		for j:=i+1 ; j<len(arr) ; j++ {
			if arr[j]<arr[i] {
				temp := arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}
	}

	fmt.Println(arr)
}
