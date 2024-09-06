package double_pointer

import (
	"fmt"
	"math"
	"testing"
)

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], k = 3
输出：true

示例 2：
输入：nums = [1,0,1,1], k = 1
输出：true

示例 3：
输入：nums = [1,2,3,1,2,3], k = 2
输出：false
*/
func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k)) // true

	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k)) // true

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(containsNearbyDuplicate(nums, k)) // false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int, 0)

	for i, v := range nums {
		idx, ok := dict[v]
		if ok && int(math.Abs(float64(i-idx))) <= k {
			return true
		}
		dict[v] = i
	}

	return false
}
