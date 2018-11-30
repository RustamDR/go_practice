package main

import (
	"fmt"
	"unicode/utf8"
)

func checkBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func repeatedSubstringPattern(s string) bool {
	bytesString := []byte(s)

	for i := range s {

		if i < 1 && utf8.RuneCountInString(s) > 2 {
			continue
		}
		template, check := bytesString[:i+1], bytesString[i+1:]
		lenTmp := len(template)

		for chunk := 0; chunk < len(check)/lenTmp; chunk += 1 {
			if checkBytes(template, check[chunk*lenTmp:(chunk+1)*lenTmp]) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("bbaab"))
}
