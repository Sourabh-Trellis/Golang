package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go printNumbers(&wg)

	wg.Wait()
}
