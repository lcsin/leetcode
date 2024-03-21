package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	val := 2
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, val), nums)
}

func removeElement(nums []int, val int) int {
	cnt := 0
	for _, num := range nums {
		if num != val {
			nums[cnt] = num
			cnt++
		}
	}
	return cnt
}
