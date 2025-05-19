package double_pointer

import "testing"

/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

提示：
2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers 按 非递减顺序 排列
-1000 <= target <= 1000

仅存在一个有效答案
*/
func TestTwoSum2(t *testing.T) {
	t.Log(twoSum2([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum2([]int{2, 3, 4}, 6))
	t.Log(twoSum2([]int{-1, 0}, -1))
	t.Log(twoSum2([]int{-5, -4, -1, 13}, 12))
}

/*
解题思路：
1. 数组是有序递增的，使用相向双指针计算两边的两数之和sum
2. 如果sum大于target，说明需要移动右指针
3. 如果sum小于target，说明需要移动左指针
4. 直到sum等于target，说明找到了两数之和等于target的元素
*/
func twoSum2(numbers []int, target int) []int {
	var (
		left  = 0
		right = len(numbers) - 1
	)

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
