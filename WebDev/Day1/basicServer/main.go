package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandleFunc)
	fmt.Println("---")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server")
	} else {
		fmt.Println("Server started successfully")
	}
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hey ,hello")
	
}
