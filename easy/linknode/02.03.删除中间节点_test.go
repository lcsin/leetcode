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

	DeleteNode(head)
	printNode(head)
}

func DeleteNode(node *ListNode) {
	if node == nil || node.Next == nil || node.Next.Next == nil {
		return
	}

	for n := node; n != nil; n = n.Next {
		if n.Val == node.Val {

		}
	}
}
