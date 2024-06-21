package main

import (
	// "fmt"
)

type Quiz struct {
	QuestionMap map[string]string 
	Score int
}

func main()  {

	questionMap, err := ReadCSV("problems.csv")
	if err != nil {
		panic(err)
	}
	
	quiz := &Quiz {
		QuestionMap : questionMap,
		Score: 0,
	}


	startRepl(quiz)

}

