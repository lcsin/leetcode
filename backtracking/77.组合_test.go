package backtracking

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
}

var path []int     // 存储每次递归的结果值
var result [][]int // 存储最终的返回值

func combine(n int, k int) [][]int {
	path, result = make([]int, 0, 2), make([][]int, 0)

	backtrackingCombine(n, k, 1) // 回溯函数
	return result
}

func backtrackingCombine(n, k, start int) {
	// 结束递归的条件: 当path数组长度为2时，说明这次递归已经获得组合结果
	if len(path) == k {
		temp := make([]int, k) // 因为path切片共用一个底层数组,所以需要拷贝到一个中间切片在添加到result
		copy(temp, path)       // 将path切片的值拷贝到一个中间切片
		result = append(result, temp)
		return
	}

	for i := start; i <= n; i++ {
		//if n - i + 1 < k - len(path) {  // 剪枝
		//	break
		//}
		path = append(path, i)         // 处理组合节点
		backtrackingCombine(n, k, i+1) // 递归
		path = path[:len(path)-1]      // 回溯撤销处理组合节点
	}
}
