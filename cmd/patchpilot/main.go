package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/pojithakarunathilake99/patchpilot/internal/store"
)

func main() {
	ctx := context.Background()

	db, err := store.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
