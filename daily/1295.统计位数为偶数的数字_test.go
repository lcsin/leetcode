package daily

import (
	"fmt"
	"testing"
)

/*
给你一个整数数组 nums，请你返回其中包含 偶数 个数位的数字的个数。

示例 1：
输入：nums = [12,345,2,6,7896]
输出：2
解释：
12 是 2 位数字（位数为偶数）
345 是 3 位数字（位数为奇数）
2 是 1 位数字（位数为奇数）
6 是 1 位数字 位数为奇数）
7896 是 4 位数字（位数为偶数）
因此只有 12 和 7896 是位数为偶数的数字

示例 2：
输入：nums = [555,901,482,1771]
输出：1
解释：
只有 1771 是位数为偶数的数字。

提示：
1 <= nums.length <= 500
1 <= nums[i] <= 105
*/
func TestFindNumbers(t *testing.T) {
	t.Log(findNumbers([]int{12, 345, 2, 6, 7896}))
	t.Log(findNumbers([]int{555, 901, 482, 1771}))
}

func findNumbers(nums []int) int {
	var count int
	for _, num := range nums {
		s := fmt.Sprintf("%d", num)
		if len(s)%2 == 0 {
			count++
		}
	}
	return count
}

func checkNumbers(n int) bool {
	var count int
	for {
		n = n / 10
		count++
		if n == 0 {
			break
		}
	}

	return count%2 == 0
}
