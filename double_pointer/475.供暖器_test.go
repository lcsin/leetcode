package double_pointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

在加热器的加热半径范围内的每个房屋都可以获得供暖。

现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。

注意：所有供暖器 heaters 都遵循你的半径标准，加热的半径也一样。

示例 1:
输入: houses = [1,2,3], heaters = [2]
输出: 1
解释: 仅在位置 2 上有一个供暖器。如果我们将加热半径设为 1，那么所有房屋就都能得到供暖。

示例 2:
输入: houses = [1,2,3,4], heaters = [1,4]
输出: 1
解释: 在位置 1, 4 上有两个供暖器。我们需要将加热半径设为 1，这样所有房屋就都能得到供暖。

示例 3：
输入：houses = [1,5], heaters = [2]
输出：3
*/
func TestFindRadius(t *testing.T) {
	house := []int{1, 2, 3}
	heaters := []int{2}
	fmt.Println(findRadius(house, heaters)) // 1

	house = []int{1, 2, 3, 4}
	heaters = []int{1, 4}
	fmt.Println(findRadius(house, heaters)) // 1

	house = []int{1, 5}
	heaters = []int{2}
	fmt.Println(findRadius(house, heaters)) // 3

	house = []int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}
	heaters = []int{499, 500, 501}
	fmt.Println(findRadius(house, heaters)) // 498

	house = []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}
	heaters = []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}
	fmt.Println(findRadius(house, heaters)) // 161834419

	house = []int{1, 2, 3, 4, 5, 4}
	heaters = []int{1, 2, 3, 4, 5, 4}
	fmt.Println(findRadius(house, heaters)) // 0
}

/*
两个序列的双指针：

算法思路：获取每个房子距离最短的供暖器的距离，再从这些距离里面取最大值，即为最终结果

算法实现：
1. 先对houses和heaters排序（因为测试用例存在无序的情况）
2. p1指针指向houses，p2指针指向heaters
3. 计算p1指向的值和p2、p2+1指向的值的距离：r1和r2
4. 当r1>=r2时，说明r1这个供暖器可能不是最优解，需要移动p2指针，指向下一个供暖器
5. 否则说明r1这个供暖器是最优解，此时需要判断这个距离是否大于之前记录的距离，大则记录，否则不记录，并移动p1指针，指向下一个房子
6. 当所有的房子遍历完后，返回记录的距离即为最优解

注意：移动p2指针时，需要判断是否越界，即供暖器存在一个的情况
*/
func findRadius(houses []int, heaters []int) int {
	var (
		p1     = 0
		p2     = 0
		r1     = 0
		r2     = 0
		radius = 0
	)

	sort.Ints(houses)
	sort.Ints(heaters)
	for p1 < len(houses) {
		r1 = abs(houses[p1] - heaters[p2])
		if p2+1 < len(heaters) {
			r2 = abs(houses[p1] - heaters[p2+1])
		} else {
			r2 = r1
		}

		if r1 >= r2 && p2+1 < len(heaters) {
			p2++
		} else {
			p1++
			m := min(r1, r2)
			if m > radius {
				radius = m
			}
		}
	}

	return radius
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
