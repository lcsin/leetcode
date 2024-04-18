package backtracking

import (
	"fmt"
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
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treepaths []string
var treepath string

func binaryTreePaths(root *TreeNode) []string {
	treepaths = make([]string, 0)
	backtrackingBinaryTreePaths(root)
	return treepaths
}

func backtrackingBinaryTreePaths(root *TreeNode) {
	if root == nil {
		treepaths = append(treepaths, treepath)
		return
	}
	treepath += fmt.Sprintf("%v->", root.Val)
	backtrackingBinaryTreePaths(root.Left)
	backtrackingBinaryTreePaths(root.Right)
}
