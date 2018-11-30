package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	bytesString, checkInd, upperIndOfTemplate, byteInTemplate, startCheck := []byte(s), 0, 1, false, false

	for i := 1; i < len(bytesString); {

		byteInTemplate = bytesString[i] == bytesString[checkInd]
		if byteInTemplate {

			startCheck = checkInd == 0
			if startCheck {
				upperIndOfTemplate = i
			}

			checkInd++ // shifting checkInd to next in template
			i++        // shifting byte in string
			continue
		}

		if checkInd == 0 {
			i++ // just shifting to next byte
		} else {
			i = upperIndOfTemplate + 1 // try to expand previous template
		}

		checkInd = 0 // starting again checking the pattern
	}

	return checkInd > 0 && checkInd%upperIndOfTemplate == 0
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcababcab"))
}
