package handlers

import (
	"concurency/functions"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	functions.CreateFile("abc")
	fmt.Fprintf(w, "File created successfully")
}
