package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	j := make(map[rune]interface{}, len(J))
	for _, s := range J {
		j[s] = struct {
		}{}
	}

	count := 0
	for _, s := range S {
		if _, ok := j[s]; ok {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "ssdfdf"))
}
