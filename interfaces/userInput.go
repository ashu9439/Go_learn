package main

import (
	"fmt"
	"time"
)

// Person struct with Name and DOB fields
type Person struct {
	Name string
	DOB  time.Time
}

func main() {
	var name string
	var year, month, day int

	// Get user input for the name
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	// Get user input for the date of birth
	fmt.Print("Enter year of birth (YYYY): ")
	fmt.Scanln(&year)
	fmt.Print("Enter month of birth (MM): ")
	fmt.Scanln(&month)
	fmt.Print("Enter day of birth (DD): ")
	fmt.Scanln(&day)

	// Create a time.Time object for the date of birth
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Create a Person instance with the input data
	person := Person{
		Name: name,
		DOB:  dob,
	}

	// Print the struct
	fmt.Printf("Person: %+v\n", person)
}
