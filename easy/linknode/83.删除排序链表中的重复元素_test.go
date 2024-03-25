package linknode

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}}
	printNode(head)

	newHead := deleteDuplicates(head)
	printNode(newHead)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	for item := head; item.Next != nil; {
		if item.Val == item.Next.Val {
			item.Next = item.Next.Next
		} else {
			item = item.Next
		}
	}

	return head
}
