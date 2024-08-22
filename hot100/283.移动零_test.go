package hot100

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
	nums := []int{73348, 66106, -85359, 53996, 18849, -6590, -53381, -86613, -41065, 83457, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 一边记录0的个数，一边把不为0的元素往前移，然后再末尾补0
func moveZeroes(nums []int) {
	idx := 0
	for _, v := range nums {
		if v != 0 {
			nums[idx] = v
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}
