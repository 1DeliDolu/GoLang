package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ExportJSON() {
	people := []Person{
		{"Peter", 44, "male"},
		{"Lois", 43, "female"},
	}

	// Create the full path to data.json
	dataPath := filepath.Join(".", "data.json")

	files, err := os.Create(dataPath)
	if err != nil {
		panic(err)
	}
	defer files.Close()

	encoder := json.NewEncoder(files)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(people); err != nil {
		panic(err)
	}
	fmt.Println("Data successfully exported to data.json")
}
