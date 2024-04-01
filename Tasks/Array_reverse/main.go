package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of array")
	fmt.Scan(&size)
	arr := make([]int,size)

	fmt.Printf("Enter %v elements in the array",size)
	for i:=0;i<size;i++{
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)
	var j = size-1

	for i := 0; i < size/2; i++ {
		temp := arr[i]
		arr[i]=arr[j]
		arr[j]=temp
		j--
	}
	fmt.Println(arr)
}
