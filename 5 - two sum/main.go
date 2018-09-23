package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i, v := range nums {
		if previousIndex, ok := mp[target-v]; ok {
			return []int{previousIndex, i}
		}

		mp[v] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
