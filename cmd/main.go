package main

import (
	"log"

	"github.com/killerekt/grpc-prac/config"
	"github.com/killerekt/grpc-prac/internal/db"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to get values from the env")
	}

	err = db.DBInit(cfg)
	if err != nil {
		log.Fatalf("Failed to Initialise the DB")
	}
}
