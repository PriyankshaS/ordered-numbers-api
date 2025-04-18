package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ordered-numbers-api/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiToken := os.Getenv("API_TOKEN")
    if apiToken == "" {
        log.Fatal("API_TOKEN not found in .env")
    }

    // Create a new router using Gorilla Mux
    router := mux.NewRouter()

    // Create a handler for the /ordered-numbers endpoint
    numbersHandler := handlers.NewNumbersHandler(apiToken, os.Getenv("API_URL"))
    router.HandleFunc("/ordered-numbers", numbersHandler.GetOrderedNo).Methods("GET")

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port
    }
    fmt.Printf("Server listening on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}