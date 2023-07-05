package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emivespa/mutant/prisma/db"
)

func mutantHandler(client *db.PrismaClient, ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	// We wrap the function so it has access to the Prisma client and the ctx.
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println("Failed to decode JSON:", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		isMutantDna, err := isMutant(req.Dna)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if isMutantDna {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}

		dnaBytes, err := json.Marshal(req.Dna)
		if err != nil {
			return
		}
		dna := string(dnaBytes)

		// We process the mutant candidate in a goroutine,
		// to return 200/403 as soon as we know whether it's a match.
		opCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		go processMutantCandidate(client, opCtx, cancel, dna, isMutantDna)
	}
}

func processMutantCandidate(client *db.PrismaClient, ctx context.Context, cancel context.CancelFunc, dna string, isMutantDna bool) error {
	defer cancel()
	_, err := client.MutantCandidate.UpsertOne(
		db.MutantCandidate.Dna.Equals(dna),
	).Create(
		db.MutantCandidate.Dna.Set(dna),
		db.MutantCandidate.IsMutant.Set(isMutantDna),
	).Update().Exec(ctx)
	return err
}
