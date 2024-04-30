package main

import "fmt"

func main() {

	var ch chan string = make(chan string)
	ch <- "sdsd"
	v := <-ch
	fmt.Println(v)

}
