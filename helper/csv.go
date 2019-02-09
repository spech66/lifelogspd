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
