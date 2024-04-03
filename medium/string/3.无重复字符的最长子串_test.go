package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "asjrgapa"
	fmt.Println(lengthOfLongestSubstringByIteration(s))
	fmt.Println(lengthOfLongestSubstringByDoublePointer(s))
}

func lengthOfLongestSubstringByIteration(s string) int {
	m := make(map[string]struct{})
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		for j := i + 1; j < len(s); j++ {
			if !strings.ContainsAny(str, string(s[j])) {
				str += string(s[j])
			} else {
				m[str] = struct{}{}
				break
			}
		}
		if len(str) > 0 {
			m[str] = struct{}{}
		}
	}

	var max int
	for k, _ := range m {
		if len(k) > max {
			max = len(k)
		}
	}

	return max
}

// todo 有问题待修改
func lengthOfLongestSubstringByDoublePointer(s string) int {
	m := make(map[uint8]struct{})
	left, right := 0, 0

	for right < len(s) {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = struct{}{}
		} else {
			delete(m, s[left])
			left++
		}
		right++
	}

	return len(m)
}
