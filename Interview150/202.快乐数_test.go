package double_pointer

import (
	"fmt"
	"testing"
)

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：
输入：n = 19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

示例 2：
输入：n = 2
输出：false
*/
func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

/*
思路：
1. 每次将n各个位置上的平方和sum存放到map中
2. 当sum的值为1时，返回true
3. 当sum的值已经存在于map中时，返回false
*/
func isHappy(n int) bool {
	dict := make(map[int]bool)
	for {
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		if dict[sum] {
			return false
		}

		dict[sum] = true
		n = sum
	}
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		mod := n % 10
		sum += mod * mod
		n /= 10
	}
	return sum
}
