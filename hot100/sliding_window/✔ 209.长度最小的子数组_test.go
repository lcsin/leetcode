package sliding_window

import "testing"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

提示：
1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104

进阶：
如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
*/
func TestMinSubArrayLen(t *testing.T) {
	t.Log(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))               // 2
	t.Log(minSubArrayLen(4, []int{1, 4, 4}))                        // 1
	t.Log(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))        // 0
	t.Log(minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8})) // 2
	t.Log(minSubArrayLen(6, []int{10, 2, 3}))                       // 1
}

/*
滑动窗口：在窗口内保持sum>=target的同时，缩小范围，得到最小值
1. 初始状态下，left、right指针都从0开始向右移动
2. left指向满足条件的子数组的左端点，主要用于缩小窗口范围；right指向满足条件的子数组的右端点，主要用于扩大窗口范围
3. 当sum-nums[left]>=target时，尝试移动left指针，缩小窗口范围，直到sum小于target，那么表示找到了最小值，进行更新
4. 直到nums数组枚举结束，返回答案
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, sum, left := n+1, 0, 0
	// 右指针指向数组右断点，扩大窗口范围
	for right, x := range nums {
		sum += x
		// 满足和大于等于target时，移动left指针，尝试缩小窗口范围
		for sum-nums[left] >= target {
			sum -= nums[left]
			left++
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}
