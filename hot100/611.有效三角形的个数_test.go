package hot100

import (
	"sort"
	"testing"
)

/*
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。

示例 1:
输入: nums = [2,2,3,4]
输出: 3
解释:有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3

示例 2:
输入: nums = [4,2,3,4]
输出: 4

提示:
1 <= nums.length <= 1000
0 <= nums[i] <= 1000
*/

func TestName(t *testing.T) {
	t.Log(triangleNumber([]int{2, 2, 3, 4}))                // 3
	t.Log(triangleNumber([]int{4, 2, 3, 4}))                // 4
	t.Log(triangleNumber([]int{48, 66, 61, 46, 94, 75}))    // 19
	t.Log(triangleNumber([]int{24, 3, 82, 22, 35, 84, 19})) // 10
}

/*
1. 三角形成立的条件：任意两边之和大于第三边
2. 先对nums数组从小到大排序，这样可以利用数组的有序性
3. 从左边开始遍历时，只要左边的两条边之和大于最右边，就可以判断中间的几条边都可以组成三角形
3. 依次对nums数组的每条边进行枚举遍历，第三条边从右边开始枚举遍历，当满足两边之和大于第三边时
4. 可以直接计算中间的条数，然后继续遍历第二条边，以此类推，直到遍历完三条边
*/
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var (
		res int
		n   = len(nums)
	)

	for i := 0; i < n-2; i++ {
		x := nums[i]
		for j := i + 1; j < n-1; j++ {
			var temp int
			for k := len(nums) - 1; k > j; k-- {
				if x+nums[j] > nums[k] {
					temp = k
					break
				}
			}
			res += max(temp-j, 0)
		}
	}

	return res
}
