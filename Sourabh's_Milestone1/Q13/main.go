package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) {

	for i := 1; i < 6; i++ {
		ch <- i
	}
	wg.Done()
}

func reciever(ch chan int, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println(num)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	wg.Add(1)
	go sender(ch, &wg)

	wg.Add(1)
	go reciever(ch, &wg)

	wg.Wait()

}
