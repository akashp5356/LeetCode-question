package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]   // stores the maximum sum found so far
	currentSum := nums[0] // stores the maximum sum of subarray ending at current index

	for i := 1; i < len(nums); i++ {
		// Either extend the current subarray OR start a new one at nums[i]
		currentSum = max(nums[i], currentSum+nums[i])

		// Update the global max
		maxSoFar = max(maxSoFar, currentSum)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
