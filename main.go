package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go running on Kubernetes!")
		fmt.Fprintf(w, "Result: %d\n", Add(10, 20))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "healthy")
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("Go application running securely on port 8080")
	log.Fatal(server.ListenAndServe())
}
