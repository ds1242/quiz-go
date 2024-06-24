package main

import (
	// "fmt"
	"flag"
)

type Quiz struct {
	QuestionData []questionData
	Score int
}


func main()  {
	pathVar := flag.String("path", "problems.csv", "path to csv file with Quiz information")
	flag.Parse()
	questionInfo, err := ReadCSV(*pathVar)
	if err != nil {
		panic(err)
	}
	
	quiz := &Quiz {
		QuestionData : questionInfo,
		Score: 0,
	}
	startRepl(quiz)
}

