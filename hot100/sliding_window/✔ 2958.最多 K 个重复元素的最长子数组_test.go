package sliding_window

import "testing"

/*
给你一个整数数组 nums 和一个整数 k 。

一个元素 x 在数组中的 频率 指的是它在数组中的出现次数。

如果一个数组中所有元素的频率都 小于等于 k ，那么我们称这个数组是 好 数组。

请你返回 nums 中 最长好 子数组的长度。

子数组 指的是一个数组中一段连续非空的元素序列。

示例 1：
输入：nums = [1,2,3,1,2,3,1,2], k = 2
输出：6
解释：最长好子数组是 [1,2,3,1,2,3] ，值 1 ，2 和 3 在子数组中的频率都没有超过 k = 2 。[2,3,1,2,3,1] 和 [3,1,2,3,1,2] 也是好子数组。
最长好子数组的长度为 6 。

示例 2：
输入：nums = [1,2,1,2,1,2,1,2], k = 1
输出：2
解释：最长好子数组是 [1,2] ，值 1 和 2 在子数组中的频率都没有超过 k = 1 。[2,1] 也是好子数组。
最长好子数组的长度为 2 。

示例 3：
输入：nums = [5,5,5,5,5,5,5], k = 4
输出：4
解释：最长好子数组是 [5,5,5,5] ，值 5 在子数组中的频率没有超过 k = 4 。
最长好子数组的长度为 4 。

提示：
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= nums.length
*/
func TestMaxSubarrayLength(t *testing.T) {
	t.Log(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)) // 6
	t.Log(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)) // 2
	t.Log(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))    // 4
	t.Log(maxSubarrayLength([]int{1, 4, 4, 3}, 1))             // 2
	t.Log(maxSubarrayLength([]int{1, 1, 1, 3}, 2))             // 3
}

/*
1. 使用一个map来记录nums中元素出现的次数
2。每次移动右指针时，判断map中是否存在次数大于k的元素
3. 如果存在，左端点元素出现的次数-1，并移动左指针，直到left>right或者map中不存在次数大于k的时候
4. 否则更新最大值
*/
func maxSubarrayLength(nums []int, k int) int {
	if len(nums) < k {
		return len(nums)
	}

	cnt := make(map[int]int)
	var ans, left int
	for right, num := range nums {
		cnt[num]++
		for left < right && cnt[num] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
