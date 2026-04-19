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
		port = "3000" // changed default from 8080 to 3000 for my local setup
	}

	log.Printf("ds2api %s (%s) built at %s, server on port %s", version, commit, date, port)

	server := NewServer()
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
