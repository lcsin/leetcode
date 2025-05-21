package sliding_window

import (
	"strings"
	"testing"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb")) // 3
	t.Log(lengthOfLongestSubstring("bbbbb"))    // 1
	t.Log(lengthOfLongestSubstring("pwwkew"))   // 3
	t.Log(lengthOfLongestSubstring("asjrgapa")) // 6
	t.Log(lengthOfLongestSubstring("dvdf"))     // 3
}

/*
1. 因为s由英文字母、数字、符号和空格组成，所以可以借助数组来协助判断子串中是否存在重复的字符
2. 每次移动窗口的右端点时，即对数组中该字符的次数+1
3. 当数组中某个字符出现的次数大于1时，说明此时子串中存在重复字符，此时需要缩小窗口，移动左端点，移动之前需要对左端点字母出现的次数-1
4. 每次移动右端点，更新结果的最大值
*/
func lengthOfLongestSubstring(s string) int {
	var (
		ans, left int
		cnt       = [128]int{}
	)

	for right, c := range s {
		cnt[c]++
		// 说明存在重复字符
		for cnt[c] > 1 {
			cnt[s[left]]-- // 移除窗口左端点字母
			left++         // 缩小窗口
		}
		ans = max(ans, right-left+1) // 更新窗口长度最大值
	}

	return ans
}

// 不借助数组来构造窗口的方法，通过左右指针的移动来构造窗口
func lengthOfLongestSubstring2(s string) int {
	var (
		ans, left int
		str       string
	)

	for right := 0; right < len(s); right++ {
		for left < right {
			if strings.ContainsAny(str, string(s[right])) {
				ans = max(ans, len(str))
				left++
				str = s[left:right]
			} else {
				break
			}
		}
		str = s[left : right+1]
	}

	return max(ans, len(str))
}
