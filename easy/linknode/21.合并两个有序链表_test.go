package linknode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	printNode(node1)
	printNode(node2)

	newNode := mergeTwoLists(node1, node2)
	printNode(newNode)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	l := p

	for {
		if list1 == nil {
			p.Next = list2
			break
		}
		if list2 == nil {
			p.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	return l.Next
}
