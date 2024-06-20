package main

import (
	// "fmt"
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
	
	startRepl(questionMap)

}

