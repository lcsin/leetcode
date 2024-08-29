package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums), nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums), nums)
}

/*
题目解析：
1. 因为nums数组是递增排列的，所以几个相同的元素一定是相邻的
2. 我们可以使用双指针，p0指针指向不重复元素序列的结尾，p1指针指向待验证序列的开头
3. 每次遍历移动p1指针，判断p0和p1的值是否相等，相等则继续移动p1，不相等则移动p0，然后将p1的值赋值给p0
4. 当p1遍历结束后，p0+1即为不重复元素的个数，同时nums前p0+1个元素一定是唯一的

算法实现：
1. 指针p0和p1同时指向nums数组下标0的位置
2. 每次遍历判断nums[p0]是否等于nums[p1]，相等则只移动p1，不相等则移动p0，并将p1的值赋值给p0
3. 遍历结束后，返回p0+1（因为p0是从下标0开始，所以返回个数需要+1）
*/
func removeDuplicates(nums []int) int {
	var (
		p0 = 0
		p1 = 0
	)

	for p1 < len(nums) {
		if nums[p0] != nums[p1] {
			p0++
			nums[p0] = nums[p1]
		}
		p1++
	}

	return p0 + 1
}
