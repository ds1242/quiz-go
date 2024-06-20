package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(questionMap map[string]string) {
	for key, val := range questionMap {
		fmt.Printf("%s = %s\n", key, val)
	}
}