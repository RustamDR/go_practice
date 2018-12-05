package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkDoublicates(field int) map[string][]string {

	if field > 6 || field < 0 {
		panic("Field is out of range")
	}

	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter, splited, line := make(map[string][]string), make([]string, 7), ""
	result := make(map[string][]string)

	for scanner.Scan() {

		line = scanner.Text()
		splited = strings.Split(line, ":")
		uuid, userName := splited[field], splited[0]

		if _, ok := counter[uuid]; !ok {
			counter[uuid] = make([]string, 0)
		}

		counter[uuid] = append(counter[uuid], userName)
	}

	for uuid, users := range counter {
		if len(users) > 1 {
			result[uuid] = users
		}
	}

	return result
}

func main() {
	fmt.Println(checkDoublicates(4))
}
