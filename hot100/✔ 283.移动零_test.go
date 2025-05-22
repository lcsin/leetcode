package hot100

import "testing"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

提示:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

进阶：你能尽量减少完成的操作次数吗？
*/

func TestMoveZeroes(t *testing.T) {
	nums := []int{1, 0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

/*
官方题解：
使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

注意到以下性质：
- 左指针左边均为非零数；
- 右指针左边直到左指针处均为零。
因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
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

/*
1. 遍历nums数组，把不等于0的元素往前移动，最后再填充0
*/
func moveZeroes2(nums []int) {
	var idx int
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
