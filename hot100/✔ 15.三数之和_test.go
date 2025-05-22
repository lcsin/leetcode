package hot100

import (
	"sort"
	"testing"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。

示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。

提示：
3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func TestThreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{0, 1, 1}))
	t.Log(threeSum([]int{0, 0, 0}))
	t.Log(threeSum([]int{-2, 0, 1, 1, 2}))
}

/*
解题思路：
1. 先将数组排序
2. 然后将三数之和转变为两数之和，转变的思路是循环枚举第一个数
3. 最后使用相向双指针遍历所有可能的答案
4. 题目要求答案中不可以包含重复的三元组，所以还要考虑去重的情况（可以考虑用map去重）
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		x := nums[i]
		j := i + 1
		k := len(nums) - 1
		// 去重
		if i > 0 && x == nums[i-1] {
			continue
		}
		// 优化，如果最小的三个数之和大于0，说明后面不存在和等于0的情况
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 优化，如果x和最大的两个数之和小于0，说明x不存在和等于0的情况
		if x+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}
		for j < k {
			sum := x + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				// 去重
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
