package double_pointer

import (
	"fmt"
	"strings"
	"testing"
)

/*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

示例 1：
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。

示例 2：
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。

示例 3：
输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。由于空字符串正着反着读都一样，所以是回文串。
*/
func TestIsPalindrome(t *testing.T) {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
	//fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("A man, a plan, a canae: Panama"))
}

/*
ASCII码：
0~9：48~57
a~z：97~122
A~Z：65~90
*/
func isPalindrome(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}

	var (
		p0 = 0
		p1 = len(s) - 1
	)

	for p0 != p1 {
		s1 := string(s[p0])
		s2 := string(s[p1])

		if !checkASCII(p0) {
			p0++
			continue
		}
		if !checkASCII(p1) {
			p1--
			continue
		}

		if strings.EqualFold(s1, s2) {
			p0++
			p1--
		} else {
			return false
		}
	}

	return true
}

func checkASCII(ascii int) bool {
	if (ascii >= 48 && ascii <= 57) || (ascii >= 97 && ascii <= 122) || (ascii >= 65 && ascii <= 90) {
		return true
	}
	return false
}
