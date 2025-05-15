package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//go:generate easyjson -all main.go

// Person represents a person's information
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address string   `json:"address,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

func main() {
	// Create a person
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
		Tags:    []string{"employee", "developer"},
	}

	// Marshal to JSON using standard library
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("JSON:", string(data))

	// Unmarshal from JSON
	var newPerson Person
	err = json.Unmarshal(data, &newPerson)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Unmarshaled: %+v\n", newPerson)

	// Call placeholder function
	DemoEasyJSONUsage()

	// Call Kubernetes dependencies demo
	DemoClient()
	DemoController()
}
