package hot100

import (
	"sort"
	"testing"
)

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

提示：
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/

func TestFourSum(t *testing.T) {
	t.Log(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	t.Log(fourSum([]int{2, 2, 2, 2, 2}, 8))
	t.Log(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}

// 解题思路与三数之和一致
// 去重方式：
// - 使用map去重最简单
// - 通过比较上一个值是否相等
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		res [][]int
		//dict = make(map[string]bool)
	)

	for a := 0; a < n-3; a++ {
		x := nums[a]
		// 去重，a是第一次枚举，要保证a大于0，且x的值与上一个值相等则跳过
		if a > 0 && x == nums[a-1] {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			y := nums[b]
			c := b + 1
			d := n - 1
			// 去重，b=a+1，要保证b-1大于a，且y的值与上一个值相等则跳过
			if b-1 > a && y == nums[b-1] {
				continue
			}
			for c < d {
				sum := x + y + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{x, y, nums[c], nums[d]})
					// 通过map去重
					//key := fmt.Sprintf("%v%v%v%v", x, y, nums[c], nums[d])
					//if !dict[key] {
					//	res = append(res, []int{x, y, nums[c], nums[d]})
					//	dict[key] = true
					//}
					//c++
					//d--

					// 通过比较前一个值去重
					c++
					// c是左指针，要保证c小于d，且c是往右移动的，所以上一个值是c-1，当与上一个值相等时则跳过
					for c < d && nums[c] == nums[c-1] {
						c++
					}
					d--
					// d是右指针，要保证d大于c，且d是往左移动的，所以上一个值是d+1，当与上一个值相等时则跳过
					for d > c && nums[d] == nums[d+1] {
						d--
					}
				} else if sum > target {
					d--
				} else {
					c++
				}
			}
		}
	}

	return res
}
