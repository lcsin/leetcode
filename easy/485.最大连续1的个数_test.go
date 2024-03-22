package easy

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var result, count int

	for _, num := range nums {
		switch num {
		case 0:
			count = 0
		case 1:
			count++
		}
		if result < count {
			result = count
		}
	}

	return result
}
