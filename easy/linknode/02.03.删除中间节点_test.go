package linknode

import "testing"

func TestDeleteNode(t *testing.T) {
	head := GenListNode([]int{4, 5, 1, 9})
	PrintNode(head)

	deleteNode(head.Next)
	PrintNode(head)
}

func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}
