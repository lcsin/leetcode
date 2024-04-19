package backtracking

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestBinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 37,
		Left: &TreeNode{
			Val:  34,
			Left: nil,
			Right: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   48,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treepaths []string
var treenodes []string

func binaryTreePaths(root *TreeNode) []string {
	// 重置全局变量
	defer func() {
		treepaths = make([]string, 0)
		treenodes = make([]string, 0)
	}()

	treepaths = make([]string, 0)
	treenodes = make([]string, 0)
	backtrackingBinaryTreePaths(root)
	return treepaths
}

func backtrackingBinaryTreePaths(root *TreeNode) {
	treenodes = append(treenodes, fmt.Sprintf("%v", root.Val))
	// 遍历到子节点，添加结果
	if root != nil && root.Left == nil && root.Right == nil {
		treepaths = append(treepaths, strings.Join(treenodes, ""))
		return
	}
	// 有左子树，递归遍历左子树
	if root.Left != nil {
		treenodes = append(treenodes, "->")
		backtrackingBinaryTreePaths(root.Left)
		treenodes = treenodes[:len(treenodes)-2]
	}
	// 有右子树，递归遍历右子树
	if root.Right != nil {
		treenodes = append(treenodes, "->")
		backtrackingBinaryTreePaths(root.Right)
		treenodes = treenodes[:len(treenodes)-2]
	}
}
