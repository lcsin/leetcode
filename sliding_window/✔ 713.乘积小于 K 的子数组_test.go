package sliding_window

import "testing"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

示例 1：
输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2]、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。

示例 2：
输入：nums = [1,2,3], k = 0
输出：0

提示:
1 <= nums.length <= 3 * 104
1 <= nums[i] <= 1000
0 <= k <= 106
*/
func TestNumSubarrayProductLessThanK(t *testing.T) {
	t.Log(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	t.Log(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

/*
思路大体上和【209.长度最小的子数组】类似，不同的是，该题不需要找到最小值
1. 初始状态下，left、right指针都从0开始向右移动
2. left指向满足条件的子数组的左端点，主要用于缩小窗口范围；right指向满足条件的子数组的右端点，主要用于扩大窗口范围
3. 当乘积大于等于k时，尝试移动left指针，缩小窗口范围，直到乘积小于k，那么代表找到了答案
4. 直到nums数组枚举结束，返回答案
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var (
		ans, left int
		res       = 1
	)
	// right指向满足条件的子数组的右端点
	for right, x := range nums {
		res *= x
		// 满足乘积大于等于k，移动left指针，尝试缩小窗口范围
		for res >= k {
			res /= nums[left]
			left++
		}
		// 当[left, right]的乘积小于k时，期间的元素也都小于k
		// 所以答案的个数就为 right-left+1，+1是因为下标从0开始
		ans += right - left + 1
	}

	return ans
}
