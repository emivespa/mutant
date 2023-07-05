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
		var countResult []map[string]interface{}
		opCtx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()

		err = client.Prisma.QueryRaw(
			"SELECT isMutant, count(*) FROM `MutantCandidate` GROUP BY isMutant",
		).Exec(opCtx, &countResult)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(countResult)

		// True comes before false:
		mutantCountStr := countResult[0]["count(*)"].(string)
		humanCountStr := countResult[1]["count(*)"].(string)
		mutantCount, _ := strconv.Atoi(mutantCountStr)
		humanCount, _ := strconv.Atoi(humanCountStr)

		ratio := -1.0
		if humanCount != 0 {
			ratio = float64(mutantCount) / float64(humanCount)
		}

		data := map[string]interface{}{
			"count_mutant_dna": mutantCountStr,
			"count_human_dna":  humanCountStr,
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
