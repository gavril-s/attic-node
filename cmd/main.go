package main

import (
	"log"
	"net/http"

	"github.com/gavril-s/attic-node/internal/app"
	"github.com/gavril-s/attic-node/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/read", app.Read)
	mux.HandleFunc("/write", app.Write)

	log.Printf("Starting the server at %s\n", cfg.Address())
	err = http.ListenAndServe(cfg.Address(), mux)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
