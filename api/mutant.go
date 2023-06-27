package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/emivespa/mutant/mutant"
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		isMutant := mutant.IsMutant(req.Dna)

		if isMutant {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}

		dnaBytes, err := json.Marshal(req.Dna)
		if err != nil {
			return
		}
		dnaString := string(dnaBytes)

		go func() {
			_, err := client.MutantCandidate.FindFirst(
				db.MutantCandidate.DnaString.Equals(string(dnaString)),
			).Exec(ctx)
			if errors.Is(err, db.ErrNotFound) {
				createdMutantCandidate, err := client.MutantCandidate.CreateOne(
					db.MutantCandidate.DnaString.Set(string(dnaString)),
					db.MutantCandidate.IsMutant.Set(isMutant),
				).Exec(ctx)
				if err != nil {
					log.Println("error creating row:", err)
				}
				res, _ := json.Marshal(createdMutantCandidate)
				log.Println("ereated row:", string(res))
			} else if err != nil {
				log.Printf("error occurred: %s", err)
			} else {
				log.Println("row exists")
			}
		}()
	}
}
