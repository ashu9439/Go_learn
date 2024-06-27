package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Post struct to hold the response from the API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// do function to make the API call with headers and return the response
func do(url string) (*Post, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating the request: %w", err)
	}

	// Set custom headers
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	req.Header.Set("Accept", "application/json")

	// Make the HTTP request using the Do method
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making the GET request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %w", err)
	}

	// Unmarshal the JSON response into the Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling the response body: %w", err)
	}

	return &post, nil
}

func main() {
	// API URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Call the do function to get the Post data
	post, err := do(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response
	fmt.Printf("Post: %+v\n", post)
}
