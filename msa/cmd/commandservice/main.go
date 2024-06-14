package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, CommandService!!!"))
	})

	log.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadEnv() error {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env.commandservice.dev"
	}
	return godotenv.Load(envFile)
}
