package main

import "fmt"

var morseAlphabet = map[rune]string{'a': ".-", 'b': "-...", 'c': "-.-.", 'd': "-..", 'e': ".", 'f': "..-.", 'g': "--.",
	'h': "....", 'i': "..", 'j': ".---", 'k': "-.-", 'l': ".-..", 'm': "--", 'n': "-.", 'o': "---", 'p': ".--.", 'q': "--.-",
	'r': ".-.", 's': "...", 't': "-", 'u': "..-", 'v': "...-", 'w': ".--", 'x': "-..-", 'y': "-.--", 'z': "--.."}

func uniqueMorseRepresentations(words []string) int {

	unique, morse := make(map[string]struct{}), ""

	for _, word := range words {
		morse = ""
		for _, w := range word {
			morse += morseAlphabet[w]
		}
		if morse == "" {
			continue
		}
		if _, ok := unique[morse]; !ok {
			unique[morse] = struct{}{}
		}
	}

	return len(unique)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
