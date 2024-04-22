package main

import "fmt"

func sumarray(arr [5]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func main() {

	arr := [5]int{}
	fmt.Println("Enter 5 numbers")
	for i := 0; i < len(arr); i++ {
		fmt.Println("Enter Number :")
		var num int
		fmt.Scan(&num)
		arr[i] = num
	}

	sum := sumarray(arr)
	fmt.Println("Sum of elements is", sum)

}
