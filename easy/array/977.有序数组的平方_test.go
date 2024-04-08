package array

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	//nums := []int{-7, -3, 2, 3, 11}
	//nums := []int{-99, -98, 1, 2, 3, 4, 5, 6}
	//nums := []int{1}
	//nums := []int{-5, -3, -2, -1}
	//nums := []int{-2, 0}
	fmt.Println(sortedSquaresByMyDoublePointer(nums))
	fmt.Println(sortedSquaresByDoublePointer(nums))
}

func sortedSquaresByDoublePointer(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}

func sortedSquaresByMyDoublePointer(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	target := make([]int, 0)
	left, right := -1, -1
	// 左右指针找到负数和正数的下标
	for i, num := range nums {
		if num < 0 {
			left = i
		} else if num >= 0 {
			right = i
			break
		}
	}

	// 双指针从中间向两边依次比较大小添加到新数组
	for left >= 0 && (right >= 0 && right < len(nums)) {
		leftVal, rightVal := nums[left]*nums[left], nums[right]*nums[right]

		if leftVal > rightVal {
			target = append(target, rightVal)
			right++
		} else {
			target = append(target, leftVal)
			left--
		}
	}

	// 右指针添加完毕，添加左指针数据
	if left >= 0 || right == -1 {
		for i := left; i >= 0; i-- {
			target = append(target, nums[i]*nums[i])
		}
	}
	// 左指针添加完毕，添加右指针数据
	if right >= len(nums) || left == -1 {
		for i := right; i < len(nums); i++ {
			target = append(target, nums[i]*nums[i])
		}
	}

	return target
}
