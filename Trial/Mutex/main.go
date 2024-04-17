package main

import (
	"fmt"
	"sync"
)

var counter int

var mutex sync.Mutex

func increment(id int) {
	mutex.Lock()
	defer mutex.Unlock() // Ensures unlock even in case of panics
	counter++
	fmt.Println("Goroutine", id, "counter:", counter)
}

func main() {
	for i := 0; i < 10; i++ {
		go increment(i)
		
	}
	fmt.Scanln() // Wait for goroutines to finish
}
