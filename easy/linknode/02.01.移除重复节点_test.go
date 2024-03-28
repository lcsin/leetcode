package linknode

import "testing"

func TestRemoveDuplicateNodes(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 3, 2, 1})
	PrintNode(head)

	newHead := removeDuplicateNodes(head)
	PrintNode(newHead)
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	m := make(map[int]struct{})
	m[head.Val] = struct{}{}

	for h := head; h.Next != nil; {
		if _, ok := m[h.Next.Val]; ok {
			h.Next = h.Next.Next
		} else {
			m[h.Next.Val] = struct{}{}
			h = h.Next
		}
	}

	return head
}
