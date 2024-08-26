package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个非负整数数组 nums，  nums 中一半整数是 奇数 ，一半整数是 偶数 。

对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。

你可以返回 任何满足上述条件的数组作为答案 。

示例 1：
输入：nums = [4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。

示例 2：
输入：nums = [2,3]
输出：[2,3]
*/
func TestSortArrayByParityII(t *testing.T) {
	nums1 := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums1))

	nums2 := []int{2, 3}
	fmt.Println(sortArrayByParityII(nums2))
}

/*
同向双指针：
1. 设置两个指针，一个指向偶数下标0，另一个指向奇数下标1，再设置一个尾指针，指向数组最后一个元素
2. 判断该元素是偶数还是奇数，如果是偶数就和偶数指针交换位置，如果是奇数就和奇数指针交换位置，并将对应的指针往下一个位置移动
3. 直到指针越界，就说明已经将元素排好序了
*/
func sortArrayByParityII(nums []int) []int {
	var (
		p1   = 0
		p2   = 1
		tail = len(nums) - 1
	)

	for {
		if p1 >= len(nums) || p2 >= len(nums) {
			break
		}

		tmp := nums[tail]
		if nums[tail]%2 == 0 {
			nums[tail] = nums[p1]
			nums[p1] = tmp
			p1 += 2
		} else {
			nums[tail] = nums[p2]
			nums[p2] = tmp
			p2 += 2
		}
	}

	return nums
}
