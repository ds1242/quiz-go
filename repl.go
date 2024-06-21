package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(quiz *Quiz) {
	reader := bufio.NewScanner(os.Stdin)
	quizLength := len(quiz.QuestionData)
	for key, val := range quiz.QuestionData {
		fmt.Printf("Question #%d %s = ", key, val.Question)
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Print("no answer provided, moving to the next question\n")
			continue
		}

		answer := words[0]

		if answer == val.Answer {
			quiz.Score++
		}	
		
	}
	fmt.Printf("Your score is %d / %d\n", quiz.Score, quizLength)
	fmt.Println("Thanks for playing!!")
	
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}