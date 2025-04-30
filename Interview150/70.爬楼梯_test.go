package double_pointer

import (
	"testing"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(1))
	t.Log(climbStairs2(1))

	t.Log(climbStairs(2))
	t.Log(climbStairs2(2))

	t.Log(climbStairs(4))
	t.Log(climbStairs2(4))

	t.Log(climbStairs(8))
	t.Log(climbStairs2(8))

	t.Log(climbStairs(45))
	t.Log(climbStairs2(45))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	mem := make([]int, n+1, n+1)
	mem[1] = 1
	mem[2] = 2

	return climb(n, mem)
}

func climb(n int, mem []int) int {
	if mem[n] != 0 {
		return mem[n]
	}
	mem[n] = climb(n-1, mem) + climb(n-2, mem)
	return mem[n]
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n+1, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
