package array

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}

func minSubArrayLen(target int, nums []int) int {
	var ans, sum int
	left, right := 0, 0
	for left < len(nums) {
		if sum < target && right < len(nums)-1 {
			sum += nums[right]
			right++
		} else {
			ans = right - left
			sum -= nums[left]
			left++
		}
	}

	return ans
}
