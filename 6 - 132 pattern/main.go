package main

import "fmt"

func find132pattern(A []int) [][3]int {
	result, temp := make([][3]int, 0), make([][2]int, 0)

	for i, v := range A {
		for ind, t := range temp {
			if t[1] > -1 && A[t[0]] < v && v < A[t[1]] {
				result = append(result, [...]int{A[t[0]], A[t[1]], v})
			}

			if v-A[t[0]] > 1 && i < len(A)-1 {
				if t[1] == -1 {
					temp[ind] = [...]int{t[0], i}
				} else {
					temp = append(temp, [...]int{t[0], i})
				}
			}
		}

		temp = append(temp, [...]int{i, -1})
	}

	return result
}

func main() {
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
}
