package easy

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//num := removeDuplicates(nums)
	//num := removeDuplicatesByDoublePointer(nums)
	num := removeDuplicatesByDoublePointer2(nums)
	fmt.Println(num, nums[:num])
}

func removeDuplicatesByDoublePointer2(nums []int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[slow] != nums[i] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
}

func removeDuplicatesByDoublePointer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fast, slow := 1, 1
	for i := fast; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	cnt := 0
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = i - cnt
		} else {
			cnt++
		}
	}

	for k, v := range m {
		nums[v] = k
	}

	return len(m)
}
