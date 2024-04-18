package backtracking

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations(""))
}

var letters []string
var letter string
var keymap = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	letters = make([]string, 0)
	if len(digits) == 0 {
		return letters
	}
	backtrackingLetterCombinations(digits, 0)
	return letters
}

func backtrackingLetterCombinations(digits string, index int) {
	if len(letter) == len(digits) {
		letters = append(letters, letter)
		return
	}

	idx, _ := strconv.Atoi(string(digits[index]))
	strs := keymap[idx-2]
	for i := 0; i < len(strs); i++ {
		letter += string(strs[i])
		backtrackingLetterCombinations(digits, index+1)
		letter = letter[:len(letter)-1]
	}
}
