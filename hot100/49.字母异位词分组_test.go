package hot100

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/
func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
}

// 将字符串数组的元素根据字典排序，将排序后的值作为map的key，将排序前的值作为map的value添加到数组中
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[string][]string, 0)

	for _, v := range strs {
		mp[sortStr(v)] = append(mp[sortStr(v)], v)
	}
	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

func sortStr(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
