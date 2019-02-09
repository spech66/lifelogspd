package helper

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadAllDataFromCSV reads all fields to interface
func ReadAllDataFromCSV(filename string) [][]string {
	fmt.Println("Read all data from", filename)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	return lines
}

// SaveDataToCSV saves all fields to CVS
func SaveDataToCSV(filename string, data []string) {
	fmt.Println("Save data to", filename)

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = ';'
	w.Write(data)
	w.Flush()
}
