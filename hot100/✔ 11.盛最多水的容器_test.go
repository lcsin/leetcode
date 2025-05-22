package hot100

import (
	"testing"
)

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1

提示：
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func TestMaxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	t.Log(maxArea([]int{1, 1}))                      // 1
}

/*
解题思路：
1. 取数组的两边作为容器的两条高，要取面积最大值，只能往中间枚举
2. 往中间枚举时，长度再减小，意味着高要变大，才能得到更大的面积
3. 所以可以考虑使用相向指针指向两条高，每次移动两条高中，值较小的那一条，并计算面积，更新面积的最大值
*/
func maxArea(height []int) int {
	var (
		left  = 0
		right = len(height) - 1
		area  = 0
	)

	for left < right {
		res := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		area = max(area, res)
	}

	return area
}
