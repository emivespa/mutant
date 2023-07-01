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
		dnaString := string(dnaBytes)

		opCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		go func(ctx context.Context, cancel context.CancelFunc) {
			defer cancel()
			upsertedMutantCandidate, err := client.MutantCandidate.UpsertOne(
				db.MutantCandidate.DnaString.Equals(string(dnaString)),
			).Create(
				db.MutantCandidate.DnaString.Set(string(dnaString)),
				db.MutantCandidate.IsMutant.Set(isMutantDna),
			).Update().Exec(ctx)
			if err != nil {
				log.Printf("error occurred: %s", err)
			} else {
				log.Println(upsertedMutantCandidate)
			}
		}(opCtx, cancel)
	}
}
