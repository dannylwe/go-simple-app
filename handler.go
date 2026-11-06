package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
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
	ctx := r.Context()
	client = NewFlagClient()
	flags, err := client.GetEnvironmentFlags(ctx)
	if err != nil {
		http.Error(w, "failed to get feature flag", http.StatusPreconditionFailed)
		return
	}

	statusCheckEnabled, err := flags.IsFeatureEnabled("status_check")
	if err != nil {
		http.Error(w, "failed to get feature flag", http.StatusPreconditionFailed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	if statusCheckEnabled {
		fmt.Fprintf(w, CheckStatusV2())
		return
	}
	fmt.Fprintf(w, CheckStatusV1())
}

func clientTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(callClient()))
}

func callClient()[]byte {
	// Create a Resty client with a timeout
	client := resty.New().SetTimeout(10 * time.Second)

	// Make the GET request
	resp, err := client.R().Get("http://172.35.8.85:9085/api/v1/phoenix/client/doKeyExchange")
	if err != nil {
		fmt.Printf("Error making HTTP call: %v\n", err)
		return []byte{}
	}

	// Resty automatically handles reading and closing the body
	fmt.Printf("Response Status: %s\n", resp.Status())
	fmt.Printf("Response Body: %s\n", resp.String())
	return resp.Body()
}
