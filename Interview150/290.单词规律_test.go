package double_pointer

import (
	"fmt"
	"strings"
	"testing"
)

/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:
输入: pattern = "abba", s = "dog cat cat dog"
输出: true

示例 2:
输入:pattern = "abba", s = "dog cat cat fish"
输出: false

示例 3:
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
*/
func TestWordPattern(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "doc cat cat fish")) // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // false
	fmt.Println(wordPattern("aabb", "dog dog dog dog"))  // false
}

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	dict1 := make(map[string]string)
	dict2 := make(map[string]string)

	for i, str := range strs {
		key := string(pattern[i])
		_, ok1 := dict1[key]
		_, ok2 := dict2[str]

		if !ok1 && !ok2 {
			dict1[key] = str
			dict2[str] = key
		} else {
			if dict1[key] != str || dict2[str] != key {
				return false
			}
		}
	}

	return true
}
