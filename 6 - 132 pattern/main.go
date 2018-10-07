package main

import "fmt"

func find132pattern(A []int) bool {
	templates132, m, template := make([][]int, 0), make(map[int]struct{}), make([]int, 2)

	for i, currentPeak := range A {
		equalsPrevToCur := i > 0 && A[i] == A[i-1]
		if equalsPrevToCur {
			continue
		}

		for ind, t := range templates132 {
			isCurValueIn132 := len(t) == 2 && t[0] < currentPeak && currentPeak < t[1]
			if isCurValueIn132 {
				return true
			}

			foundPartOf132 := currentPeak-t[0] > 1 && i < len(A)-1
			if foundPartOf132 {
				template[0] = t[0]
				template[1] = currentPeak
				if len(t) == 1 || t[1] < currentPeak {
					templates132[ind] = template
				}
			}
		}

		if _, wasValueIn132 := m[currentPeak]; !wasValueIn132 {
			templates132 = append(templates132, []int{currentPeak})
			m[currentPeak] = struct{}{}
		}
	}

	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
