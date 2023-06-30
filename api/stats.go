package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/emivespa/mutant/prisma/db"
)

func statsHandler(client *db.PrismaClient, ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	// We wrap the function so it has access to the client and the ctx.
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var err error
		var mutantCountResult []map[string]string
		var humanCountResult []map[string]string
		opCtx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()

		err = client.Prisma.QueryRaw(
			"SELECT count(*) FROM `MutantCandidate` WHERE isMutant = true LIMIT 1",
		).Exec(opCtx, &mutantCountResult)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = client.Prisma.QueryRaw(
			"SELECT count(*) FROM `MutantCandidate` WHERE isMutant = false LIMIT 1",
		).Exec(opCtx, &humanCountResult)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mutantCount, _ := strconv.Atoi(mutantCountResult[0]["count(*)"])
		humanCount, _ := strconv.Atoi(humanCountResult[0]["count(*)"])

		ratio := -1.0
		if humanCount != 0 {
			ratio = float64(mutantCount) / float64(humanCount)
		}

		data := map[string]interface{}{
			"count_mutant_dna": mutantCount,
			"count_human_dna":  humanCount,
			"ratio":            ratio,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
