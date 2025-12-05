package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if idx, found := m[target-num]; found {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

// This runs in O(n) time and uses O(n) space.
