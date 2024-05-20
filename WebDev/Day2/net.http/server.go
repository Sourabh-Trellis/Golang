package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	defer http.ListenAndServe(":8000", mux)

	mux.Handle("/", &homeHandler{})
	mux.Handle("/about", &aboutHandler{})

	// mux.HandleFunc()
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(202)
	w.Write([]byte("This is Home page."))
}

type aboutHandler struct{}

func (h *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is about page."))
}
