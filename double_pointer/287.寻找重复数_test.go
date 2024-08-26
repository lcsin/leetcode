package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：
输入：nums = [1,3,4,2,2]
输出：2

示例 2：
输入：nums = [3,1,3,4,2]
输出：3

示例 3 :
输入：nums = [3,3,3,3,3]
输出：3
*/
func TestFindDuplicate(t *testing.T) {
	nums1 := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums1))

	nums2 := []int{3, 1, 4, 3, 2}
	fmt.Println(findDuplicate(nums2))

	nums3 := []int{3, 3, 3, 3, 3}
	fmt.Println(findDuplicate(nums3))

	nums4 := []int{1, 3, 4, 2, 1}
	fmt.Println(findDuplicate(nums4))

	nums5 := []int{18, 13, 14, 17, 9, 19, 7, 17, 4, 6, 17, 5, 11, 10, 2, 15, 8, 12, 16, 17}
	fmt.Println(findDuplicate(nums5))
}

/*
快慢指针：
1. 将数组元素作为链表的指针指向的下一个元素，因为至少存在一个重复的元素，那么这个链表就一定会构成一个环
2. 通过快慢指针，就一定可以找到构成闭环的这个节点，即这个重复的元素

算法实现：
1. 慢指针一次走一步，快指针一次走两部
2. 当快慢第一次相遇时，将快指针从0开始并一次走一步，慢指针继续一次走一步
3. 当快慢第二次相遇时，他们的值一定相等，即快慢指针就是重复的元素
*/
func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
