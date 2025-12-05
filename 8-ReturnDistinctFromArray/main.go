package main

import "fmt"

func unique(arr []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
func main() {
	fmt.Println("Distinct elements", unique([]int{1, 2, 3, 4, 1, 2}))
}
