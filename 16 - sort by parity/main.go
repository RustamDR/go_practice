package main

import "fmt"

func sortArrayByParity(A []int) []int {

	odd, even := make([]int, 0), make([]int, 0)
	for _, n := range A {
		if n%2 == 0 {
			even = append(even, n)
			continue
		}
		odd = append(odd, n)
	}

	return append(even, odd...)
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
