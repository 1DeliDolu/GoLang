package main

import "fmt"

// Person struct with JSON field tags
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	fmt.Println("Starting JSON operations...")

	// Export JSON data
	ExportJSON()

	// Import JSON data
	ImportJSON()

	fmt.Println("JSON operations completed.")
}
