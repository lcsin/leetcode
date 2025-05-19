package hot100

import "testing"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]


提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
*/

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{3, 2, 4}, 6))
	t.Log(twoSum([]int{3, 3}, 6))
}

/*
解题思路：
1. 使用暴力破解的时间复杂度是O(n2)，存在时间超限的可能，要缩小时间复杂度，可以使用map
2. 每次遍历数组都将元素的值和下标记录到map中，并判断map中是否存在与该元素之和等于target的值
3. 如果存在，说明找到了两数之和等于target的元素，返回他们的下标
*/
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := dict[target-v]; ok {
			return []int{i, j}
		}
		dict[v] = i
	}
	return nil
}
