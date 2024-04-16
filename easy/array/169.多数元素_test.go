package array

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	return majorityElementByMoleVote(nums)
	//return majorityElementByMap(nums)
}

// 摩尔投票法（抵消法）
func majorityElementByMoleVote(nums []int) int {
	var major, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
		}

		if nums[i] == major {
			count++
		} else {
			count--
		}
	}

	if count > 0 {
		return major
	}
	return -1
}

func majorityElementByMap(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	n := len(nums) / 2
	for k, v := range m {
		if v > n {
			return k
		}
	}

	return 0
}
