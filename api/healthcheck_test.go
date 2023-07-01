package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	tests := []struct {
		method     string
		statusCode int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusMethodNotAllowed},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		healthcheckHandler(recorder, req)

		if recorder.Code != test.statusCode {
			t.Errorf("for %s expected %d and got %d", test.method, test.statusCode, recorder.Code)
		}
	}
}
