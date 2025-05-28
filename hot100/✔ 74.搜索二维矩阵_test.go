package hot100

import (
	"sort"
	"testing"
)

/*
给你一个满足下述两条属性的 m x n 整数矩阵：
  - 每行中的整数从左到右按非严格递增顺序排列。
  - 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

示例 1：

	[1, 3, 5, 7]
	[10, 11, 16, 20]
	[23, 30, 34, 60]

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：

	[1, 3, 5, 7]
	[10, 11, 16, 20]
	[23, 30, 34, 60]

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10^4 <= matrix[i][j], target <= 10^4
*/
func TestSearchMatrix(t *testing.T) {
	t.Log(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	t.Log(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

/*
遍历二维数组，中的数组，通过二分查找判断target是否存在
*/
func searchMatrix(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		idx := sort.SearchInts(arr, target)
		if idx < len(arr) && arr[idx] == target {
			return true
		}
	}
	return false
}
