package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	healthcheckHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
