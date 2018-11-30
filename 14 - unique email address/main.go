package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {

	unique, spl, domain, name, email := make(map[string]interface{}, 10), make([]string, 2), "", "", ""

	for _, e := range emails {
		// 1. split by @
		spl = strings.Split(e, "@")
		if len(spl) < 2 {
			continue
		}
		domain = spl[1]

		// 2. split by +
		mainName := strings.Split(spl[0], "+")

		// 3. remove .
		name = strings.Replace(mainName[0], ".", "", -1)

		// 4. combine name@domain
		email = name + "@" + domain

		// 5. check if not add
		if _, ok := unique[email]; !ok {
			unique[email] = struct {
			}{}
		}
	}

	return len(unique)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
