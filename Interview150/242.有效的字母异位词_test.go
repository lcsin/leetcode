package double_pointer

import "testing"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

提示:
1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
*/
func TestIsAnagram(t *testing.T) {
	t.Log(isAnagram("anagram", "nagaram")) // true
	t.Log(isAnagram("rat", "car"))         // false
	t.Log(isAnagram("a", "ab"))            // false
	t.Log(isAnagram("ac", "aa"))           // false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict1 := make(map[int32]int)
	dict2 := make(map[int32]int)

	for _, v := range s {
		dict1[v]++
	}
	for _, v := range t {
		dict2[v]++
	}

	for k, v := range dict1 {
		if dict2[k] != v {
			return false
		}
	}

	return true
}
