package array

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left // 获取中点所在下标
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { // 目标在左边
			right = mid - 1
		} else { // 目标在右边
			left = mid + 1
		}
	}

	return -1
}
