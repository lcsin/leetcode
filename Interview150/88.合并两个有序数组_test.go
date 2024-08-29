package double_pointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

示例 2：
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。

示例 3：
输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
*/
func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

/*
直接合并两个数组，然后排序
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func TestMergeByDoublePointer(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 4, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)

	nums1 = []int{4, 0, 0, 0, 0, 0}
	nums2 = []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
	fmt.Println(nums1)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}

/*
双指针遍历不同序列：
1. p1,p2指针分别指向num1和nums2，再用一个新的数组nums3存放合并后的序列
2. 每次比较p1和p2的值，把小的一方放到nums3里，并移动指针
3. 当p1或p2有一方大于等于m或n时，循环终止，此时p1或p2有一方数组已经合并完
4. 判断p1和p2的值是否小于m或n，如果成立，从p1或p2开始将剩余的元素合并到nums3即可
*/
func mergeByDoublePointer(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
	}

	var (
		p1    = 0
		p2    = 0
		nums3 = make([]int, 0, m+n)
	)

	for p1 < m && p2 < n {
		if nums1[p1] > nums2[p2] {
			nums3 = append(nums3, nums2[p2])
			p2++
		} else {
			nums3 = append(nums3, nums1[p1])
			p1++
		}
	}

	if p1 < m {
		nums3 = append(nums3, nums1[p1:]...)
	}
	if p2 < n {
		nums3 = append(nums3, nums2[p2:]...)
	}

	copy(nums1, nums3)
}
