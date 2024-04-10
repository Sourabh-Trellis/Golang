package main

import "fmt"

func main() {
	arr := [7]int{}

	fmt.Println("Enter 7 elements in array.")

	for i := 0; i < 7; i++ {

		var n int
		fmt.Println("Enter element :")
		fmt.Scan(&n)
		arr[i] = n
	}
	var max int
	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	// max := maxVal(arr)
	fmt.Println("Max value :", max)
}

// func maxVal(arr [7]int) int {
// 	max := 0
// 	for _, val := range arr {
// 		if max < val {
// 			max = val
// 		}
// 	}
// 	return max
// }
