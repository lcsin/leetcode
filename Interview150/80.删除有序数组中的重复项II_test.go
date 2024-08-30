package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。
*/
func TestRemoveDuplicatesII(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicatesII(nums), nums)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicatesII(nums), nums)
}

/*
思路：
1. 本题和【26.删除有序数组中的重复项】类似，区别在于允许元素重复出现两次
2. 和【26.删除有序数组中的重复项】一样，序列是有序且递增的，这表示重复的元素肯定位于相邻的位置
3. 因此p0指针同样指向已经处理完的位置，p1指针指向待处理的位置
4. 不同的是，判断交换的条件从nums[p0]!=nums[p1]变成nums[p0-2]!=nums[p1]（因为需要保证允许前两个元素可以重复）
5. 当遍历结束后，p0即为新数组的长度

实现：
1. 根据题意，当数组长度小于等于2时，可以直接返回。因此p0和p1指针可以直接从下标2开始
2. 遍历数组，当nums[p0-2]!=nums[p1]时，p1和p0交换位置，此时移动p0；否则只移动p1
3. 当遍历结束后，p0即为新数组的长度，直接返回即可
*/
func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var (
		p0 = 2
		p1 = 2
	)

	for p1 < len(nums) {
		if nums[p0-2] != nums[p1] {
			nums[p0] = nums[p1]
			p0++
		}
		p1++
	}

	return p0
}
