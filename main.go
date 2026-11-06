package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MessageHandler)
	http.HandleFunc("/flaky", FlakyHandler)
	http.HandleFunc("/check", StatusCheckHandler)
	http.HandleFunc("/client", clientTest)
	fmt.Println("starting server on port 9002")
	http.ListenAndServe(":9002", nil)
}
