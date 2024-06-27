// go mod init github.com/ashu9439/Go_learn/tree/main/Api/gorillaMusk
//  go get -u github.com/gorilla/mux
// go run .

//create executable file and run it
/*
```
	go build -o myapp
	./myapp
```
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Person struct to represent data
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	// Create a new mux router
	r := mux.NewRouter()

	// Middleware: Logging middleware
	r.Use(loggingMiddleware)

	// Routes
	r.HandleFunc("/", handleRoot).Methods("GET")
	r.HandleFunc("/info", handleInfo).Methods("GET")

	// Middleware: CORS middleware
	r.Use(corsMiddleware)

	// Start the HTTP server on port 8080 with the router
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// loggingMiddleware is a middleware that logs each incoming request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware is a middleware that adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
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

// handleInfo handles requests to the "/info" endpoint
func handleInfo(w http.ResponseWriter, r *http.Request) {
	// Create a sample info response
	info := map[string]string{
		"message": "This is an info endpoint",
	}

	// Convert info map to JSON
	jsonData, err := json.Marshal(info)
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
