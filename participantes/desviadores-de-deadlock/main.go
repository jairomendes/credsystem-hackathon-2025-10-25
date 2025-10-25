package main

import (
	"context"
	"desviadores-de-deadlock/pkg/middleware"
	"desviadores-de-deadlock/pkg/openrouter"
	"desviadores-de-deadlock/pkg/service/health"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := createServer(port)

	printStartupMessages(port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func createServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", middleware.LoggingMiddleware(health.HealthHandler))

	openrouterClient := openrouter.NewClient(os.Getenv("OPENROUTER_API_KEY"))
	openrouterClient.ChatCompletion(context.Background(), "Hello, how are you?")

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func printStartupMessages(port string) {
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Health check available at http://localhost:%s/\n", port)
}
