package double_pointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。

每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

返回 承载所有人所需的最小船数 。

示例 1：
输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)

示例 2：
输入：people = [3,2,2,1], limit = 3
输出：3
解释：3 艘船分别载 (1, 2), (2) 和 (3)

示例 3：
输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)
*/
func TestNumRescueBoats(t *testing.T) {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{2, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

/*
两边向中间的双指针：
1. 先排序，排完序后，就可以从头尾两端匹配合适的体重
2. 分一下几种情况：
  - 头尾两人体重合等于limit，两人一船
  - 头尾两人体重合大于limit，右边单独一船
  - 头尾两人体重合小于limit，两人一船

3. 左指针大于右指针后，循环终止，返回结果
*/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var (
		p1    = 0
		p2    = len(people) - 1
		count = 0
		total = 0
	)

	for {
		if p1 > p2 {
			break
		}

		total = people[p1] + people[p2]
		if total <= limit {
			count++
			p1++
			p2--
		} else if total > limit {
			count++
			p2--
		}
	}

	return count
}
