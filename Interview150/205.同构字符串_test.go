package double_pointer

import (
	"fmt"
	"testing"
)

/*
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:
输入：s = "egg", t = "add"
输出：true

示例 2：
输入：s = "foo", t = "bar"
输出：false

示例 3：
输入：s = "paper", t = "title"
输出：true
*/
func TestIsIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
	fmt.Println(isIsomorphic("badc", "baba"))   // false
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dicts := make(map[int32]uint8)
	dictt := make(map[uint8]bool)
	for i, v := range s {
		tv, ok := dicts[v]
		if !ok && !dictt[t[i]] {
			dicts[v] = t[i]
			dictt[t[i]] = true
			continue
		}
		if !ok && dictt[t[i]] {
			return false
		}
		if ok && tv != t[i] {
			return false
		}
	}

	return true
}
