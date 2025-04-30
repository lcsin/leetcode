package double_pointer

import (
	"testing"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func TestLongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	t.Log(longestCommonPrefix([]string{"a"}))
	t.Log(longestCommonPrefix([]string{"ab", "a"}))
	t.Log(longestCommonPrefix([]string{"aaa", "aa", "aaa"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		var temp string
		for j, _ := range strs[i] {
			if prefix == "" || j >= len(prefix) {
				break
			}

			if strs[i][j] == prefix[j] {
				temp += string(prefix[j])
			} else {
				break
			}
		}
		prefix = temp
	}

	return prefix
}
