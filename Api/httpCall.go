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

func main() {
	// API URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Unmarshal the JSON response into the Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error unmarshaling the response body:", err)
		return
	}

	// Print the response
	fmt.Printf("Post: %+v\n", post)
}
