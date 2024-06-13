package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, CommandService!!!"))
	})

	log.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
