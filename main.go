package main

import "fmt"

var nums = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	for _, v := range nums {
		fmt.Println(v)
		nums = append(nums, 1)
	}
	fmt.Println(nums)
}

func logicalPage(pageNo, pageSize int) []int {
	if pageNo*pageSize > len(nums) {
		return nums[(pageNo-1)*pageSize:]
	} else {
		return nums[(pageNo-1)*pageSize : pageNo*pageSize]
	}
}
