package hot100

import "testing"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

func TestTrap(t *testing.T) {
	t.Log(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Log(trap2([]int{4, 2, 0, 3, 2, 5}))
}

/*
双指针：
1. 接雨水的量=两条黑柱的最小值-高，两条黑柱分别对应前缀最大值和后缀最大值
2. 使用相向双指针，分别从两边开始遍历，每次遍历都与当前的高度比较，获取前缀最大值和后缀最大值
3. 每次遍历同样取前缀最大值和后缀最大值中小的一方，减去当前高度，得到接到水的容量
4. 然后移动小的一方的指针，继续遍历
*/
func trap2(height []int) int {
	var (
		n              = len(height)
		res            int
		left           = 0
		right          = n - 1
		preMax, sufMax int
	)

	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			res += preMax - height[left]
			left++
		} else {
			res += sufMax - height[right]
			right--
		}
	}

	return res
}
