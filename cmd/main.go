package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/laudryfadian/starter-golang/internal/injector"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	server, err := injector.InitializeServer()
	if err != nil {
		log.Fatal("Failed to initialize server:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
