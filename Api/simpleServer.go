package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Person struct to represent data
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	// Define HTTP route handlers with middleware
	http.HandleFunc("/", logMiddleware(handleRoot))

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// logMiddleware is a middleware that logs each incoming request
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	}
}

// handleRoot handles requests to the root endpoint "/"
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Create a sample Person struct
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	// Convert Person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
