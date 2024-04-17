package backtracking

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

var letters []string
var letter string
var keymap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	letters = make([]string, 0)
	backtrackingLetterCombinations(digits, 0)
	return letters
}

func backtrackingLetterCombinations(digits string, index int) {
	if len(letter) == len(digits) {
		letters = append(letters, letter)
		return
	}

	for i := 0; i < len(digits); i++ {
		idx, _ := strconv.Atoi(string(digits[i]))
		letter += string(keymap[idx][index])
		backtrackingLetterCombinations(digits, index+1)
		letter = letter[:len(letter)-1]
	}
}
