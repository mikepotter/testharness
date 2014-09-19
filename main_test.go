package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	expectedBody := "Hi there, I love golang!"
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://example.com/golang", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler(recorder, req)

	if expectedBody != recorder.Body.String() {
		t.Fatalf("expected %s. Got %s", expectedBody, recorder.Body.String())
	}
}
