package easy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var arr []int

loop1:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				arr = append(arr, i, j)
				break loop1
			}
		}
	}

	return arr
}
