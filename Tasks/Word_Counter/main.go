package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var para string

	fmt.Println("welcome to word counter!!!!!!!")
	fmt.Print("type tour paragraph here : ")
	para, _ = reader.ReadString('\n')

	fmt.Println(para)
	words := strings.Split(strings.TrimSpace(para), " ")
	for _, v := range words {
		fmt.Println(v)
	}
	fmt.Println("word count is ", len(words))
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}

	for key, val := range count {
		fmt.Println(key, val)
	}

	// fmt.Println(count)
}
