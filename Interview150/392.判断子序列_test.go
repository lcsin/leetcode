package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

致谢：

特别感谢 @pbrother 添加此问题并且创建所有测试用例。

示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false
*/
func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
	fmt.Println(isSubsequence("b", "abc"))      // true
}

/*
算法实现：
1. 定义两个指针p0和p1，p0指向字符串s，p1指向字符串t
2. 每次遍历判断p0指向的字符和p1指向的字符是否相等，相等则移动p0，否则只移动p1
3. 当遍历结束后，判断p0的值：
- p0<len(s)：说明s字符串没有遍历完，s不是t的子序列，返回false
- p0>=len(s)：说明s字符串全部遍历完，s是t的子序列，返回true

注意：遍历时要求同时满足p0<len(s)、p1<len(t)
*/
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	var (
		p0 = 0
		p1 = 0
	)

	for p0 < len(s) && p1 < len(t) {
		if s[p0] == t[p1] {
			p0++
		}
		p1++
	}

	if p0 < len(s) {
		return false
	}

	return true
}
