package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}

func holaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hola mundo")
}
