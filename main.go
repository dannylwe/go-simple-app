package main

import (
	"fmt"
	"net/http"
	"os"
)

func message() string {
	return "this is a test"
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, message())
}

func flakyHandler(w http.ResponseWriter, r *http.Request) {
	appEnv := os.Getenv("APP_ENV")
	message := "success message"

	if appEnv == "prod" {
		message = "error"
	}

	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", messageHandler)
	http.HandleFunc("/flaky", flakyHandler)
	fmt.Println("starting server on port 9002")
	http.ListenAndServe(":9002", nil)
}
