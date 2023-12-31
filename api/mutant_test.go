package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/emivespa/mutant/prisma/db"
)

func TestMutantHandler(t *testing.T) {
	tests := []struct {
		method string
		code   int
		req    Request
	}{
		{"POST", http.StatusOK, Request{Dna: []string{"AAAAAA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}}},
		{"POST", http.StatusForbidden, Request{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}}},
		{"POST", http.StatusUnprocessableEntity, Request{Dna: []string{"XXXXXX", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}},
		{"GET", http.StatusMethodNotAllowed, Request{Dna: []string{"AAAAAA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}}},
	}

	for i, test := range tests {
		ctx := context.Background()
		client, mock, ensure := db.NewMock()
		defer ensure(t)

		dnaBytes, err := json.Marshal(test.req.Dna)
		if err != nil {
			t.Fatal(err)
		}
		dna := string(dnaBytes)
		isMutantDna := test.code == http.StatusOK

		reqBytes, _ := json.Marshal(test.req)
		req, err := http.NewRequest(test.method, "", bytes.NewReader(reqBytes))
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := mutantHandler(client, ctx)
		handler(recorder, req)
		if recorder.Code != test.code {
			t.Errorf("for test %d expected %d and got %d", i, test.code, recorder.Code)
		}

		// Test the DB processing:

		expected := &db.MutantCandidateModel{
			InnerMutantCandidate: db.InnerMutantCandidate{
				ID:       0,
				Dna:      dna,
				IsMutant: isMutantDna,
			},
		}

		mock.MutantCandidate.Expect(
			client.MutantCandidate.UpsertOne(
				db.MutantCandidate.Dna.Equals(dna),
			).Create(
				db.MutantCandidate.Dna.Set(dna),
				db.MutantCandidate.IsMutant.Set(isMutantDna),
			).Update(),
		).Returns(*expected)

		opCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		if err := processMutantCandidate(client, opCtx, cancel, dna, isMutantDna); err != nil {
			t.Errorf("processMutantCandidate failed")
		}
	}
}
