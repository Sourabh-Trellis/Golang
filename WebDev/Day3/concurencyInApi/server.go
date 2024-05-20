package main

import (
	handlers "concurency/Handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/try", handlers.HomeHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server.")
	} else {
		fmt.Println("server started at http://localhost:8000")
	}
}
