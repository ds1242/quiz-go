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

		fmt.Printf("%s =", key)
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Print("no answer provided, moving to the next question\n")
			continue
		}

		answer := words[0]

		// if answer == val {
		// 	quiz.Score++
		// }
		fmt.Printf("question %s", val.Question)
		fmt.Printf(answer)
		
		
	}
	
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}