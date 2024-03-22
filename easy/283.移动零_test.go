package easy

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	num := []int{0, 1, 0, 3, 12}
	moveZeroes(num)
	fmt.Println(num)
}

func moveZeroes(nums []int) {
	cnt := 0
	for i, num := range nums {
		if num == 0 {
			cnt++
		} else {
			nums[i-cnt] = num
		}
	}
	for i := 1; i <= cnt; i++ {
		nums[len(nums)-i] = 0
	}
}
