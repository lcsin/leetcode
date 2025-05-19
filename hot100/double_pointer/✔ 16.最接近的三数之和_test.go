package double_pointer

import (
	"math"
	"sort"
	"testing"
)

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

示例 1：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。

示例 2：
输入：nums = [0,0,0], target = 1
输出：0
解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。


提示：
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/

func TestThreeSumClosest(t *testing.T) {
	t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	t.Log(threeSumClosest([]int{0, 0, 0}, 1))
	t.Log(threeSumClosest([]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, 1))
}

/*
1. 先对数组nums从小到大排序，将三数之和转变为两数之和
2. 转变的方式是，从nums数组第一位开始枚举，然后使用相向双指针遍历最接近target的和
3. 计算和的结果与target的差值，遍历比较差值最小的三数，并更新值
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		res = 10000 // 通过给的条件可以知道第一次比较的结果差值一定小于10000
		arr = make([]int, 3, 3)
	)

	for i := 0; i < n-2; i++ {
		x := nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			sum := x + nums[left] + nums[right]
			if sum == target {
				return sum
			} else {
				diff := int(math.Abs(float64(sum - target)))
				if diff < res {
					res = diff
					arr = []int{i, left, right}
				}
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return nums[arr[0]] + nums[arr[1]] + nums[arr[2]]
}
