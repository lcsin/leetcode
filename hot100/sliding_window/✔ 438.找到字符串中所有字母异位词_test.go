package sliding_window

import (
	"testing"
)

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/
func TestFindAnagrams(t *testing.T) {
	t.Log(findAnagrams("cbaebabacd", "abc"))
	t.Log(findAnagrams("abab", "ab"))
}

/*
解题思路和【3.无重复字符的最长子串】类似：
1. 因为s和p仅包含小写字母，所以可以借助数组来构造一个定长的窗口，窗口长度为p的长度
2. 通过比较s数组和p数组中，字母出现的次数是否一致，可以判断s的子串是否为p的异位词子串
3. 每次滑动窗口时，都是左右指针一起滑动，保证窗口的长度等于p的长度，每次滑动都需要更新s数组中字母出现的次数，并判断和p数组是否相等
*/
func findAnagrams(s string, p string) []int {
	// s的长度小于p的长度时，s中一定不存在p的异位词的子串
	if len(s) < len(p) {
		return nil
	}

	// 计算s、p中个字母出现的次数
	// 通过比较sArr是否等于pArr来判断s的子串是否为p的异位词子串
	var sArr, pArr [26]int
	for _, c := range p {
		pArr[c-'a']++
	}

	var ans []int
	for right, c := range s {
		sArr[c-'a']++ // 移动窗口的右端点
		// 当左右端点构造的子串（窗口）长度小于p的长度时直接跳过
		// 因为无法构成p的异位词子串
		left := right - len(p) + 1
		if left < 0 {
			continue
		}
		// 当sArr和pArr相等时，说明s的子串各字母出现的次数和p相等
		// 即找到了p的异位词子串，记录左端点到答案中
		if sArr == pArr {
			ans = append(ans, left)
		}
		// 移动窗口的左端点
		sArr[s[left]-'a']--
	}

	return ans
}
