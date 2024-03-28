package linknode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node1 := GenListNode([]int{1, 2, 4})
	node2 := GenListNode([]int{1, 3, 4})
	PrintNode(node1)
	PrintNode(node2)

	newNode := mergeTwoLists(node1, node2)
	PrintNode(newNode)
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
