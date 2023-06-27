package main

import (
	"context"
	"log"
	"net/http"

	"github.com/emivespa/mutant/prisma/db"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// Establish DB connection, https://goprisma.org/docs/getting-started/quickstart
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()

	http.HandleFunc("/", healthcheckHandler)
	http.HandleFunc("/mutant", mutantHandler(client, ctx))
	http.HandleFunc("/stats", statsHandler(client, ctx))
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
