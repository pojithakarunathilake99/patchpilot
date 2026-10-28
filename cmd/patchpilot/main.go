package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/pojithakarunathilake99/patchpilot/internal/api"
	"github.com/pojithakarunathilake99/patchpilot/internal/store"
)

func main() {
	ctx := context.Background()

	db, err := store.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := api.New(db)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Routes()))
}
