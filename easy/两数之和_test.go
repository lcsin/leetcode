package easy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum1(nums, target))
}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil
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
