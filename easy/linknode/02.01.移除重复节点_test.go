package linknode

import "testing"

func TestRemoveDuplicateNodes(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	printNode(head)

	newHead := removeDuplicateNodes(head)
	printNode(newHead)
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
