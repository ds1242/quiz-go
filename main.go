package main

import (
	"encoding/csv"
	"fmt"
	// "errors"
	"os"
)

func main()  {
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

	for key, val := range questionMap {
		fmt.Printf("%s = %s\n", key, val)
	}
}