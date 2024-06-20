package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
)


func startRepl(quiz *Quiz) {
	for key, val := range quiz.QuestionMap {
		fmt.Printf("%s = %s\n", key, val)
	}
}