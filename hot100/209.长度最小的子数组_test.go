package hot100

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
滑动窗口：
1. 使用同向双向指针，从左到右遍历，当sum<target时，移动right指针
2. 当sum >= target时，停止right指针，移动left指针，缩小范围
3. 直到right等于nums长度，且sum<target时，停止遍历
*/
func minSubArrayLen(target int, nums []int) int {
	var (
		res, left, right int
		sum              = nums[left]
	)

	for {
		if sum >= target {
			x := right - left + 1
			if res == 0 {
				res = x
			} else {
				res = min(res, x)
			}

			sum -= nums[left]
			left++
		} else {
			right++
			if right < len(nums) {
				sum += nums[right]
			}
		}

		if sum < target && right == len(nums) {
			break
		}
	}

	return res
}
