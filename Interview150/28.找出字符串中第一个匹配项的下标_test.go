package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。

示例 2：
输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
*/
func TestStrStr(t *testing.T) {
	fmt.Println(strStr("sadbutsad", "sad"))     // 0
	fmt.Println(strStr("leetcode", "leeto"))    // -1
	fmt.Println(strStr("mississippi", "issip")) // 4
}

/*
思路：
1. 遍历haystack数组，当存在haystack[i]==needle[0]时，判断haystack[i:i+len(needle)]==needle
2. 如果成立直接返回i，如果不成立则继续遍历，直到遍历结束，返回-1

注意：haystack[i:i+len(needle)]有可能越界
*/
func strStr(haystack string, needle string) int {
	var p0 int
	for p0 < len(haystack) {
		if haystack[p0] == needle[0] {
			length := p0 + len(needle)
			if length > len(haystack) {
				return -1
			}
			if haystack[p0:length] == needle {
				return p0
			}
		}
		p0++
	}

	return -1
}
