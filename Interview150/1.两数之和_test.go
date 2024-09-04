package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}

/*
思路：
1. 创建一个map用于存放每次遍历的元素，key为元素的值，value为元素的下标
2. 在每次遍历的过程中，判断map中是否存在key为target-nums[i]的键值对
3. 如果存在，说明已经找到了两数之和的下标，分别为i和map[target-nums[i]]
*/
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		if v, ok := dict[target-num]; ok {
			return []int{i, v}
		}
		dict[num] = i
	}

	return nil
}
