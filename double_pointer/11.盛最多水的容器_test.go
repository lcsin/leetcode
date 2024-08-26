package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1
*/
func TestMaxArea(t *testing.T) {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height1))

	height2 := []int{1, 1}
	fmt.Println(maxArea(height2))
}

/*
两边向中间的双指针：
1. 从头尾两边计算面积，每次循环只保留高的那一边
2. 每次循环计算面积后，保留面积大的值
3. 左指针大于右指针后，循环终止，返回结果
*/
func maxArea(height []int) int {
	var (
		p1   = 0
		p2   = len(height) - 1
		area = 0
		sum  = 0
	)

	for {
		if p1 > p2 {
			break
		}

		h1 := height[p1]
		h2 := height[p2]

		if h1 > h2 {
			sum = h2 * (p2 - p1)
		} else {
			sum = h1 * (p2 - p1)
		}
		if sum > area {
			area = sum
		}

		if h1 > h2 {
			p2--
		} else {
			p1++
		}
	}

	return area
}
