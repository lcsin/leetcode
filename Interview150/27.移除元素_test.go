package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。

示例 1：
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2,_,_]
解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。

示例 2：
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3,_,_,_]
解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
注意这五个元素可以任意顺序返回。
你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
*/
func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val), nums)
}

/*
题目解析：题目要求返回不为val的元素个数k，并且要求nums数组的前k个元素不包含val，那么只要把不为val的元素往前面移动即可

算法实现：
1. 定义一个变量cnt用于记录不为val元素的个数，并且cnt同时也表示nums数组前k个不包含val元素的下标
2. 遍历nums数组，每次判断元素是否等于val，如果不等于，将该元素直接移动到cnt下标处，并且cnt++
3. 遍历结束后，cnt即为数组中不为val的元素个数，且nums数组前k个元素也不包含val
*/
func removeElement(nums []int, val int) int {
	var cnt int
	for _, num := range nums {
		if num != val {
			nums[cnt] = num
			cnt++
		}
	}
	return cnt
}
