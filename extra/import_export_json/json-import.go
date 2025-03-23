package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ImportJSON() {
	// Create the full path to data.json
	dataPath := filepath.Join(".", "data.json")

	// Read the JSON file
	fileData, err := os.Open(dataPath)
	if err != nil {
		panic(err)
	}
	defer fileData.Close()

	decoder := json.NewDecoder(fileData)
	var people []Person
	if err := decoder.Decode(&people); err != nil {
		panic(err)
	}
	fmt.Println("People imported from data.json:")
	for _, person := range people {
		fmt.Printf("  - %s, %d years old, %s\n", person.Name, person.Age, person.Gender)
	}
}
