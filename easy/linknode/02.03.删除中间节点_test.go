package linknode

import "testing"

func TestDeleteNode(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	printNode(head)

	deleteNode(head.Next)
	printNode(head)
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
