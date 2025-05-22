package binary_search

import (
	"sort"
	"testing"
)

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 10^5
-109 <= nums[i] <= 10^9
nums 是一个非递减数组
-109 <= target <= 10^9
*/
func TestSearchRange(t *testing.T) {
	t.Log(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange2([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(searchRange2([]int{}, 0))
	t.Log(searchRange2([]int{1, 2, 3}, 1))
}

/*
1. 通过二分查找找到target在nums第一次出现的下标start
2. 再通过二分查找找到target+1在nums中第一次出现的下标end，那么end-1就是target在nums中的结束位置
3. 返回[start, end]即可
4. 如果start等于nums数组长度（处理nums为空数组的情况）或者nums[start]!=target，说明数组中不存在target
*/
func searchRange(nums []int, target int) []int {
	start := sort.SearchInts(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.SearchInts(nums, target+1) - 1
	return []int{start, end}
}

// 二分查找
// 如果找到，返回target所在的下标
// 如果没找到，返回target应该插入的下标位置
func binarySearch(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		// 能够防止`mid := (left + right) / 2`溢出
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// target不在数组中，且大于数组中元素的最大值，需要在数组尾部插入
	if left == len(nums)-1 && target > nums[left] {
		return left + 1
	}

	// left就是当前数组中target的第一个目标
	// 或者target不在数组中时需要插入的下标位置
	return left
}

func searchRange2(nums []int, target int) []int {
	start := binarySearch(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := binarySearch(nums, target+1) - 1
	return []int{start, end}
}
