package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emivespa/mutant/prisma/db"
)

func TestMutantHandler(t *testing.T) {
	mockClient := &db.PrismaClient{}
	mockCtx := context.Background()

	tests := []struct {
		code int
		req  Request
	}{
		{
			http.StatusOK, Request{Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
		},
		{
			http.StatusForbidden, Request{Dna: []string{"ACACAC", "GTGTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		},
		{
			http.StatusUnprocessableEntity, Request{Dna: []string{"XXXXXX", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
		},
	}

	for i, test := range tests {
		reqBytes, _ := json.Marshal(test.req)
		req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(reqBytes))
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := mutantHandler(mockClient, mockCtx)
		handler(recorder, req)
		if recorder.Code != test.code {
			t.Errorf("for test %d expected %d and got %d", i, test.code, recorder.Code)
		}
	}
}
