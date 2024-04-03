package main

import (
	"fmt"
	"sort"
)

func main() {

	m := make(map[int]string)

	m[1] = "ford"
	m[4] = "Toyota"
	m[3] = "Ferarri"
	m[6] = "Audi"
	m[2] = "Suzuki"

	for k, v := range m {
		fmt.Printf("%v = %v\n",k,v)
	}
	fmt.Println("-------------")
	arr := []int{}

	for k, _ := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for _,val := range arr {
		fmt.Printf("%v = %v\n", val, m[val])
	}

}
