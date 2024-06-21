package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(quiz *Quiz) {
	reader := bufio.NewScanner(os.Stdin)
	
	for key, val := range quiz.QuestionData {
		quizLength := len(quiz.QuestionData)
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
		if key == quizLength {
			fmt.Printf("Your score is %d out of %d\n", quiz.Score, quizLength)
			fmt.Println("Thanks for playing")
			break
		}
		
		
	}
	
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}