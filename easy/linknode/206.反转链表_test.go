package linknode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	printNode(head)
	newHead := reverseList(head)
	printNode(newHead)

}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = newHead
		newHead = curr

		curr = next
	}

	return newHead
}
