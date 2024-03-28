package linknode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveElements(t *testing.T) {
	head := GenListNode([]int{1, 2, 6, 3, 4, 5, 6})
	PrintNode(head)

	removeElements(head, 6)
	PrintNode(head)
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Next: head}

	for temp := newHead; temp.Next != nil; {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return newHead.Next
}
