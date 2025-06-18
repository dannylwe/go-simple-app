package main

import (
	"fmt"
	"net/http"
)

func message() string {
	return "this is a test"
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, message())
}

func main() {
	http.HandleFunc("/", messageHandler)
	fmt.Println("starting server on port 9002")
	http.ListenAndServe(":9002", nil)
}
