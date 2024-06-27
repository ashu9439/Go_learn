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

// Age function calculates and returns the age of the person
func (p Person) Age() int {
	today := time.Now()
	age := today.Year() - p.DOB.Year()

	// Check if the birthday has occurred this year
	if today.YearDay() < p.DOB.YearDay() {
		age--
	}

	return age
}

// UpdateDOB function updates the DOB of the person
func (p *Person) UpdateDOB(newDOB time.Time) {
	p.DOB = newDOB
}

func main() {
	// Example date of birth: January 1, 1990
	dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Create a Person instance
	person := Person{
		Name: "John Doe",
		DOB:  dob,
	}

	// Print the person's name and age
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age())


	// Update the person's date of birth
	person.UpdateDOB( time.Date(1985, time.June, 15, 0, 0, 0, 0, time.UTC))

	// Print the person's name and age
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age())
}