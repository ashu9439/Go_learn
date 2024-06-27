package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"time"
)

// Person struct with Name and DOB fields
type Person struct {
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
}

// SavePersonToFile saves a Person struct to a file in JSON format
func SavePersonToFile(filename string, person Person) error {
	data, err := json.Marshal(person)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ReadPersonFromFile reads a Person struct from a file in JSON format
func ReadPersonFromFile(filename string) (Person, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Person{}, err
	}
	var person Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return Person{}, err
	}
	return person, nil
}

func main() {
	// Hard-coded values for name and date of birth
	name := "John Doe"
	dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Create a Person instance with the hard-coded data
	person := Person{
		Name: name,
		DOB:  dob,
	}

	// Save the Person struct to a file
	filename := "person.json"
	err := SavePersonToFile(filename, person)
	if err != nil {
		fmt.Println("Error saving person to file:", err)
		return
	}
	fmt.Println("Person saved to file:", filename)

	// Read the Person struct from the file
	readPerson, err := ReadPersonFromFile(filename)
	if err != nil {
		fmt.Println("Error reading person from file:", err)
		return
	}
	fmt.Printf("Person read from file: %+v\n", readPerson)
}
