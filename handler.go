package main

import (
	"fmt"
	"net/http"
	"os"
)

func FlakyHandler(w http.ResponseWriter, r *http.Request) {
	appEnv := os.Getenv("APP_ENV")
	message := "success message"

	if appEnv == "prod" {
		message = "error"
	}

	fmt.Fprintln(w, message)
}

func getMessage() string {
	return "this is a test"
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getMessage())
}

func StatusCheckHandler(w http.ResponseWriter, r *http.Request) {
	client = NewFlagClient()
	isEnabled, err := client.FeatureEnabled("status_check")
	if err != nil {
		http.Error(w, "failed to get feature flag", http.StatusPreconditionFailed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	if isEnabled {
		fmt.Fprintf(w, CheckStatusV2())
		return
	}
	fmt.Fprintf(w, CheckStatusV1())
}
