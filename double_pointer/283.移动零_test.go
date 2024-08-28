package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/
func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 0, 0, 4, 5}
	moveZeroes(nums)
	fmt.Println(nums)
}

/*
同向双指针：
1. p1指向已经处理好的子序列的尾部，p2指向待处理的子序列的头部
2. 当p2指向0时和p1交换位置，直到p2遍历结束
*/
func moveZeroes(nums []int) {
	var p1, p2 int
	for p2 < len(nums) {
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
		}
		p2++
	}
}
