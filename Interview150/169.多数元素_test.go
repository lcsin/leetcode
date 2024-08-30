package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
输入：nums = [3,2,3]
输出：3

示例 2：
输入：nums = [2,2,1,1,1,2,2]
输出：2
*/
func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

/*
思路：
1. 已知可以假设数组非空且数组总存在多数元素，那么经过排序后，就可以直接得出nums[len(nums)/2]即为多数元素（因为多数元素要求出现次数大于n/2）
2. 因此最简单的方法就是先对数组排序，然后直接返回len(nums)/2下标所在位置的元素即可，时间复杂度取决于你的排序算法
3. 通过map可以将时间复杂度缩短至O(n)，具体做法是遍历数组，将元素出现的次数存在map中，然后再通过map计算出现次数最多的元素

进阶：可以通过摩尔投票法将时间复杂度和空间复杂度缩短至O(n)和O(1)
*/
func majorityElement(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	var maxK, maxV int
	for k, v := range mp {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}

	return maxK
}

func TestMajorityElementByMooreVote(t *testing.T) {
	fmt.Println(majorityElementByMooreVote([]int{3, 2, 3}))
	fmt.Println(majorityElementByMooreVote([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElementByMooreVote(nums []int) int {
	var (
		candidate = 0
		count     = 0
	)

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
