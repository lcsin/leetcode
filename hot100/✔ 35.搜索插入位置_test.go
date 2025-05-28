package hot100

import "testing"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

提示:
1 <= nums.length <= 10^4
-104 <= nums[i] <= 10^4
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 10^4
*/
func TestName(t *testing.T) {
	t.Log(searchInsert([]int{1, 3, 5, 6}, 5))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 2))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 7))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 6))
}

/*
二分查找算法的变种
*/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1 // 防止溢出
		// 这里判断的顺序会影响返回的结果
		// 因为在找不到target时，要返回插入位置的下标
		// 所以要优先判断nums[mid]<target，将搜索区间向左搜索
		// 当left>right结束循环时，left恰好指向第一个大于或等于目标值的位置
		// 如果目标值存在，left就是第一个目标值的下标
		// 如果目标值不存在，left就是应该插入的位置
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
