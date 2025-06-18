package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessage(t *testing.T) {
	expected := "this is a test"
	got := message()

	if got != expected {
		t.Errorf("was expecting %s got %s", expected, got)
	}
}

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "this is a test"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
