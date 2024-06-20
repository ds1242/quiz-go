package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(quiz *Quiz) {
	for key, val := range quiz.QuestionMap {
		fmt.Printf("%s = %s\n", key, val)
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}