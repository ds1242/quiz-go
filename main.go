package main

import (
	// "fmt"
)

type Quiz struct {
	QuestionData []questionData
	Score int
}


func main()  {
	questionInfo, err := ReadCSV("problems.csv")
	if err != nil {
		panic(err)
	}
	
	quiz := &Quiz {
		QuestionData : questionInfo,
		Score: 0,
	}
	startRepl(quiz)
}

