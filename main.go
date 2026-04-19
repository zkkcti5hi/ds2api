package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Load .env file if present (ignored in production where env vars are set directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("ds2api %s (%s) built at %s", version, commit, server on port %s", port)

	server := NewServer()
	if err := server.Run" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
