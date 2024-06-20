package main

import (
	"os"
	"encoding/csv"
)

func ReadCSV(path string) (error, map[string]string) {
	questionMap := make(map[string]string)
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		questionMap[row[0]] = row[1]
	}
	return nil, questionMap
}