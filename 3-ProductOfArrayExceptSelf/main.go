package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Step 1: prefix products
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// Step 2: suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // [24 12 8 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // [0 0 9 0 0]
}
