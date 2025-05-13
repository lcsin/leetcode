package hot100

import (
	"sort"
	"testing"
)

/*
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target ，请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目。


示例 1：
输入：nums = [-1,1,2,3,1], target = 2
输出：3
解释：总共有 3 个下标对满足题目描述：
- (0, 1) ，0 < 1 且 nums[0] + nums[1] = 0 < target
- (0, 2) ，0 < 2 且 nums[0] + nums[2] = 1 < target
- (0, 4) ，0 < 4 且 nums[0] + nums[4] = 0 < target
注意 (0, 3) 不计入答案因为 nums[0] + nums[3] 不是严格小于 target 。

示例 2：
输入：nums = [-6,2,5,-2,-7,-1,3], target = -2
输出：10
解释：总共有 10 个下标对满足题目描述：
- (0, 1) ，0 < 1 且 nums[0] + nums[1] = -4 < target
- (0, 3) ，0 < 3 且 nums[0] + nums[3] = -8 < target
- (0, 4) ，0 < 4 且 nums[0] + nums[4] = -13 < target
- (0, 5) ，0 < 5 且 nums[0] + nums[5] = -7 < target
- (0, 6) ，0 < 6 且 nums[0] + nums[6] = -3 < target
- (1, 4) ，1 < 4 且 nums[1] + nums[4] = -5 < target
- (3, 4) ，3 < 4 且 nums[3] + nums[4] = -9 < target
- (3, 5) ，3 < 5 且 nums[3] + nums[5] = -3 < target
- (4, 5) ，4 < 5 且 nums[4] + nums[5] = -8 < target
- (4, 6) ，4 < 6 且 nums[4] + nums[6] = -4 < target


提示：
1 <= nums.length == n <= 50
-50 <= nums[i], target <= 50
*/

func TestCountPairs(t *testing.T) {
	t.Log(countPairs([]int{-1, 1, 2, 3, 1}, 2))
	t.Log(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
	t.Log(countPairs([]int{6, -1, 7, 4, 2, 3}, 8))

	t.Log(countPairs2([]int{-1, 1, 2, 3, 1}, 2))
	t.Log(countPairs2([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
	t.Log(countPairs2([]int{6, -1, 7, 4, 2, 3}, 8))
}

/*
1. 先对nums从小到大进行排序，这样可以利用数组的有序性，
2. 对nums进行枚举，使用双指针从两边开始遍历，
3. 如果nums[left]+nums[right]<target，说明right到left的元素都符合条件，可以直接开始下一次枚举
4. 否则移动right指针，继续判断nums[left]+nums[right]<target，直到left>=right开始下一次枚举
*/
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums)-1; i++ {
		j := len(nums) - 1
		for i < j {
			if nums[i]+nums[j] < target {
				res += max(j-i, 0)
				break
			} else {
				j--
			}
		}
	}

	return res
}

// 排序+循环+剪枝
func countPairs2(nums []int, target int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				res++
			} else {
				break
			}
		}
	}

	return res
}

// 排序+二分查找
func countPairs3(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		// sort.SearchInts: 返回待查找元素在数组中的下标或待插入位置的下标
		res += sort.SearchInts(nums[0:i], target-nums[i])
	}
	return res
}
