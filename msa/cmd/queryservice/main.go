package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, QueryService!!!"))
	})

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadEnv() error {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env.queryservice.dev"
	}
	return godotenv.Load(envFile)
}
