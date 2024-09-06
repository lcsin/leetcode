package double_pointer

import (
	"fmt"
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
}

func wordPattern(pattern string, s string) bool {
	dict := make(map[uint8]string, 0)

	var (
		p0    int
		p1    int
		space int
		str   string
	)
	for p1 < len(s) {
		if string(s[p1]) == " " || p1 == len(s)-1 {
			if p1 == len(s)-1 {
				str = s[space:len(s)]
			} else {
				str = s[space:p1]
			}

			v, ok := dict[pattern[p0]]
			if !ok {
				dict[pattern[p0]] = str
			} else {
				if str != v {
					return false
				}
			}

			space = p1 + 1
			p0++
		}

		p1++
	}

	return true
}
