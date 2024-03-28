package linknode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	PrintNode(head)
	newHead := reverseList(head)
	PrintNode(newHead)

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
