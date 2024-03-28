package linknode

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 3})
	PrintNode(head)

	newHead := deleteDuplicates(head)
	PrintNode(newHead)
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
