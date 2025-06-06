package binary_search

import (
	"sort"
	"testing"
)

/*
给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。

同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。

请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。

示例 1：
输入：spells = [5,1,3], potions = [1,2,3,4,5], success = 7
输出：[4,0,3]
解释：
- 第 0 个咒语：5 * [1,2,3,4,5] = [5,10,15,20,25] 。总共 4 个成功组合。
- 第 1 个咒语：1 * [1,2,3,4,5] = [1,2,3,4,5] 。总共 0 个成功组合。
- 第 2 个咒语：3 * [1,2,3,4,5] = [3,6,9,12,15] 。总共 3 个成功组合。
所以返回 [4,0,3] 。

示例 2：
输入：spells = [3,1,2], potions = [8,5,8], success = 16
输出：[2,0,2]
解释：
- 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
- 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
- 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
所以返回 [2,0,2] 。

提示：
n == spells.length
m == potions.length
1 <= n, m <= 10^5
1 <= spells[i], potions[i] <= 10^5
1 <= success <= 10^10
*/
func TestSuccessfulPairs(t *testing.T) {
	t.Log(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	t.Log(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
}

/*
1. 只要找到 potions[i] >= success/spells[i] 的元素，即可满足条件
2. 所以条件转换成在potions中找到potions[i]>=success/spells[i]所在元素的下标idx
3. 那么len(potions)-idx即为spells[i]和potions[i]能够成功组合的个数
4. 为了能够使用二分查找，可以先对potions数组进行排序，然后枚举spells中的元素进行查找
*/
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	pairs := make([]int, len(spells))
	for i, v := range spells {
		x := (int(success) + v - 1) / v    // 结果向上取整
		idx := sort.SearchInts(potions, x) // 二分查找
		pairs[i] = len(potions) - idx
	}

	return pairs
}
