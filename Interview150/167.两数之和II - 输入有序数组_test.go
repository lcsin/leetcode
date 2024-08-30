package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
*/
func TestName(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	target := 9
	fmt.Println(twoSum(numbers, target))

	numbers = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(numbers, target))

	numbers = []int{2, 3, 4}
	target = 6
	fmt.Println(twoSum(numbers, target))

	numbers = []int{-1, 0}
	target = -1
	fmt.Println(twoSum(numbers, target))
}

/*
思路：
1. 已知number数组递增，且一定存在两个元素之和等于target
2. 那么我们使用左右指针，从两边开始遍历求和就一定能找到答案

实现：
1. p0指针从0开始，p1指针从len(numbers)-1开始，每次循环求两者之和sum
2. 如果sum等于target，找到答案，返回p0+1和p1+1
3. 如果sum小于target，说明sum偏小，p0指针往右移动
4. 如果sum大于target，说明sum偏大，p1指针往左移动
*/
func twoSum(numbers []int, target int) []int {
	var (
		p0  = 0
		p1  = len(numbers) - 1
		sum = 0
	)

	for p0 < p1 {
		sum = numbers[p0] + numbers[p1]
		if sum == target {
			break
		} else if sum < target {
			p0++
		} else {
			p1--
		}
	}

	return []int{p0 + 1, p1 + 1}
}
