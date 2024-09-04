package double_pointer

import (
	"fmt"
	"testing"
)

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：
输入：ransomNote = "a", magazine = "b"
输出：false

示例 2：
输入：ransomNote = "aa", magazine = "ab"
输出：false

示例 3：
输入：ransomNote = "aa", magazine = "aab"
输出：true
*/
func TestCanConstruct(t *testing.T) {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	dict := make(map[int32]int, 0)
	for _, v := range magazine {
		dict[v]++
	}

	for _, v := range ransomNote {
		c, ok := dict[v]
		if !ok || c == 0 {
			return false
		}
		dict[v]--
	}

	return true
}
