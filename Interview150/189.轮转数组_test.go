package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}

/*
思路：遍历数组，将元素移动k个位置后，放到一个新的数组里面，然后再将新数组赋值给nums

实现：
1. 创建一个空数组arr，长度和nums一致
2. 每次遍历nums，就将元素移动k个位置后，再赋值给arr：arr[(i+k)%len(nums)] = num
3. 最后将arr赋值给nums：copy(nums, arr)
*/
func rotate(nums []int, k int) {
	arr := make([]int, len(nums))
	for i, num := range nums {
		arr[(i+k)%len(nums)] = num
	}
	copy(nums, arr)
}

func TestRotateBySlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotateBySlice(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotateBySlice(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100}
	k = 3
	rotateBySlice(nums, k)
	fmt.Println(nums)
}

/*
思路：
1. 通过观察可以发现，数组轮转k个元素，就相当于将k之前的元素放到k之后，k之后的元素放到k之前
2. 需要注意的是，k的大小有可能超过数组长度，所以再通过k做切片运算时，需要对数组长度取余

实现：
1. 创建一个新的数组arr，用于存放nums对k做切片运算后的两个数组
2. 根据k对nums做切片运算，并添加到arr
3. 将arr数组赋值给nums
*/
func rotateBySlice(nums []int, k int) {
	arr := make([]int, 0, len(nums))
	k = k % len(nums)
	arr = append(arr, nums[len(nums)-k:]...)
	arr = append(arr, nums[:len(nums)-k]...)
	copy(nums, arr)
}
