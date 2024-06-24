package main

import (
	"encoding/csv"
	// "fmt"
	"os"
)
type questionData struct {
	Question string
	Answer string
}

func ReadCSV(path string) ([]questionData, error) {
	questionInfo := make([]questionData, 0)
	file, err := os.Open(path)
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
		questionInfo = append(questionInfo, questionData{row[0], row[1]})
	}
	return questionInfo, nil
}