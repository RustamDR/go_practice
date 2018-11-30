package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	searchState = iota
	multiState
)

func main() {
	state, path, word := searchState, os.Args[1], `.*`+os.Args[2]+`.*`

	inSingleComment := regexp.MustCompile(`(?i)` + `^\/\/` + word)

	multiCommentStart := regexp.MustCompile(`^\/\*\*`)
	multiCommentEnd := regexp.MustCompile(`\*\*\/$`)
	inMultiComment := regexp.MustCompile(`(?i)` + word)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()

		switch {

		// Search in one line comment
		case state == searchState:
			if multiCommentStart.MatchString(txt) {
				state = multiState
			} else if inSingleComment.MatchString(txt) {
				fmt.Println(strings.TrimSpace(txt))
			}

			// If multi line comment started
			continue

		case state == multiState:
			if inMultiComment.MatchString(txt) {
				fmt.Println(strings.TrimSpace(txt))
			}
			if multiCommentEnd.MatchString(txt) {
				state = searchState
			}
		}
	}
}
