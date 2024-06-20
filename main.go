package main

import (
	"fmt"
)

type QuestionAnswer struct {
	Question string 
	Answer string
}

func main()  {

	err, questionMap := ReadCSV("problems.csv")
	if err != nil {
		panic(err)
	}
	

	for key, val := range questionMap {
		fmt.Printf("%s = %s\n", key, val)
	}
}

