package main

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		array []int
	}{
		{[]int{1, 2, 3, 4}},
		{[]int{-1, 1, 0, -3, 3}},
	}
	for _, ut := range tests {
		err := productExceptSelf(ut.array)
		if err != nil {
			fmt.Println("Error", err)
		}
	}
}
